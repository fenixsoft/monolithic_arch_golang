package payment

import (
	"github.com/fenixsoft/monolithic_arch_golang/domain"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/state"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/util"
)

type StockpileService struct {
	*state.STDCtx
}

func (s *StockpileService) decrease(productId uint, amount int) {
	stock := util.Try(domain.NewStockpile(s.Database()).GetByProductId(productId)).(*domain.Stockpile)
	stock.Frozen -= amount
	util.Try(stock.Update())
	s.Logger().Infof("库存出库，商品：%v，数量：%v\n", productId, amount)
}

func (s *StockpileService) thawed(productId uint, amount int) {
	stock := util.Try(domain.NewStockpile(s.Database()).GetByProductId(productId)).(*domain.Stockpile)
	stock.Amount += amount
	stock.Frozen -= amount
	util.Try(stock.Update())
	s.Logger().Infof("库存入库，商品：%v，数量：%v\n", productId, amount)
}

func (s *StockpileService) frozen(productId uint, amount int) {
	stock := util.Try(domain.NewStockpile(s.Database()).GetByProductId(productId)).(*domain.Stockpile)
	stock.Amount -= amount
	stock.Frozen += amount
	util.Try(stock.Update())
	s.Logger().Infof("冻结库存，商品：%v，数量：%v\n", productId, amount)
}
