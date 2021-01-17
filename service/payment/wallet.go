package payment

import (
	"errors"
	"github.com/fenixsoft/monolithic_arch_golang/domain"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/state"
)

type WalletService struct {
	*state.STDCtx
}

//账户资金减少
func (s *WalletService) decrease(accountId uint, amount float64) error {
	wallet := try(domain.NewWallet(s.Database()).GetByAccountId(accountId)).(*domain.Wallet)
	// 如果没有钱包记录（譬如新注册用户），先插入一条
	if wallet.IsNew() {
		wallet.AccountId = accountId
		wallet.Money = 0
	}
	if wallet.Money >= amount {
		wallet.Money -= amount
		try(wallet.Update())
		s.Logger().Infof("支付成功。用户余额：%v，本次消费：%v\n", wallet.Money, amount)
		return nil
	} else {
		return errors.New("用户余额不足以支付，请先充值")
	}
}

// 后面应该与库存服务一样，有类似的increase、frozen、thawed方法，由于这个是演示，就不添加了。
