// 支付应用服务
package payment

import (
	"errors"
	"github.com/fenixsoft/monolithic_arch_golang/domain"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/db"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/state"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

const defaultFrozenExpires = 2 * 60 * 1000

var (
	cache = make(map[string]*domain.Settlement)
	try   = util.Try
)

type Service struct {
	*state.STDCtx
}

func New() *Service {
	return &Service{STDCtx: state.New()}
}

func WithGinContext(ctx *gin.Context) *Service {
	return &Service{STDCtx: state.WithGinContext(ctx)}
}

// 根据结算清单的内容执行，生成对应的支付单
func (s *Service) ExecuteBySettlement(bill *domain.Settlement) *domain.Payment {
	s.ReplenishProductInformation(bill)
	payment := s.ProducePayment(bill)
	go s.setupAutoThawedTrigger(payment)
	return payment
}

// 根据结算单中货物的ID，填充货物的完整信息到结算单对象上
func (s *Service) ReplenishProductInformation(bill *domain.Settlement) {
	var ids = make([]uint, len(bill.Items))
	for k, v := range bill.Items {

		ids[k] = v.ID
	}
	r, _ := domain.NewProduct(s.Database()).FindByIDs(ids)
	for k, v := range r {
		bill.ProductMap[v.ID] = &r[k] // 不能用&v，循环中同个v地址轮询各个元素
	}
}

// 生成支付单
// 根据结算单冻结指定的货物，计算总价，生成支付单
func (s *Service) ProducePayment(bill *domain.Settlement) *domain.Payment {
	stock := &StockpileService{STDCtx: s.STDCtx}
	// 计算总价，12元固定运费，客户端写死的，这里陪着演一下，避免总价对不上
	var total float64 = 12
	for _, v := range bill.Items {
		stock.frozen(v.ID, v.Amount)
		total += float64(v.Amount) * bill.ProductMap[v.ID].Price
	}
	// 生成支付单
	payment := domain.NewPaymentWithInfo(s.Database(), s.LoginUser(), total, defaultFrozenExpires)
	try(payment.Create())
	cache[payment.PayId] = bill
	s.Logger().Infof("创建支付订单，总额：%v\n", total)
	return payment
}

// 设置支付单自动冲销解冻的触发器
// 如果在触发器超时之后，如果支付单未仍未被支付（状态是WAITING）
// 则自动执行冲销，将冻结的库存商品解冻，以便其他人可以购买，并将Payment的状态修改为ROLLBACK。
func (s *Service) setupAutoThawedTrigger(payment *domain.Payment) {
	defer delete(cache, payment.PayId)
	logger := logrus.StandardLogger()
	time.Sleep(defaultFrozenExpires * time.Millisecond)
	// 支付状态很可能已经改变，以数据库中实际状态为准
	if currentPayment, err := domain.NewPayment(db.NewDB()).Get(payment.ID); err != nil {
		// 这里出错就需要人工介入了
		logger.Warnf("支付状态变更失败，数据库中查询不到支付单(%v)出错：%v\n", payment.ID, err.Error())
	} else {
		if currentPayment.PayState == domain.Waiting {
			logger.Infof("支付单%v当前状态为：WAITING，转变为：TIMEOUT\n", payment.ID)
			s.AccomplishSettlement(domain.Timeout, payment.PayId)
		} else {
			logger.Infof("支付单%v当前状态为：%v，无需变更状态\n", payment.ID, domain.ConvPayStateToString(currentPayment.PayState))
		}
	}
}

// 根据支付状态，实际调整库存（扣减库存或者解冻）
func (s *Service) AccomplishSettlement(payState domain.PayState, payId string) {
	settlement := cache[payId]
	if settlement != nil {
		stock := &StockpileService{STDCtx: s.STDCtx}
		for _, v := range settlement.Items {
			if payState == domain.Payed {
				stock.decrease(v.ID, v.Amount)
			} else {
				// 其他状态，无论是TIMEOUT还是CANCEL，都进行解冻
				stock.thawed(v.ID, v.Amount)
			}
		}
	} else {
		s.Logger().Warnf("缓存中不存在支付单%v\n", payId)
	}
}

// 完成支付
// 立即取消解冻定时器，执行扣减库存和资金
// 意味着客户已经完成付款，这个方法在正式业务中应当作为三方支付平台的回调，而演示项目就直接由客户端发起调用了
func (s *Service) AccomplishPayment(accountId uint, payId string) error {
	payment := try(domain.NewPayment(s.Database()).GetByPayID(payId)).(*domain.Payment)
	if payment.PayState == domain.Waiting {
		payment.PayState = domain.Payed
		try(payment.Update())
		s.AccomplishSettlement(domain.Payed, payment.PayId)
		wallet := &WalletService{STDCtx: s.STDCtx}
		s.Logger().Infof("编号为%v的支付单已处理完成，等待支付\n", payId)
		return wallet.decrease(accountId, payment.TotalPrice)
	} else {
		return errors.New("当前订单不允许支付，当前状态为：" + domain.ConvPayStateToString(payment.PayState))
	}
}

// 取消支付
// 立即取消解冻定时器，执行扣减库存和资金
// 意味着客户已经完成付款，这个方法在正式业务中应当作为三方支付平台的回调，而演示项目就直接由客户端发起调用了
func (s *Service) CancelPayment(payId string) error {
	payment := try(domain.NewPayment(s.Database()).GetByPayID(payId)).(*domain.Payment)
	if payment.PayState == domain.Waiting {
		payment.PayState = domain.Cancel
		try(payment.Update())
		s.AccomplishSettlement(domain.Cancel, payment.PayId)
		s.Logger().Infof("编号为%v的支付单已被取消\n", payId)
		return nil
	} else {
		return errors.New("当前订单不允许取消，当前状态为：" + domain.ConvPayStateToString(payment.PayState))
	}
}
