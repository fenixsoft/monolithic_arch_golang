package controller

import (
	"github.com/fenixsoft/monolithic_arch_golang/domain"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/response"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/state"
	"github.com/fenixsoft/monolithic_arch_golang/service/payment"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PaymentController struct{}

func (c *PaymentController) Register(router gin.IRouter) {
	// 变更支付单据的状态
	router.PATCH("/pay/:payId", func(context *gin.Context) {
		name := state.WithGinContext(context).LoginUser()
		accountId := try(domain.NewAccount(DB(context)).GetByName(name)).(*domain.Account).ID
		payId := context.Param("payId")
		payState := try(domain.ConvStringToPayState(context.Query("state"))).(domain.PayState)
		updatePaymentState(context, payId, accountId, payState)
	})

	/**
	 * 修改支付单状态的GET方法别名
	 * 考虑到该动作要由二维码扫描来触发，只能进行GET请求，所以增加一个别名以便通过二维码调用
	 * 这个方法原本应该作为银行支付接口的回调，不控制调用权限（谁付款都行），但都认为是购买用户付的款
	 */
	router.GET("/pay/modify/:payId", func(context *gin.Context) {
		payId := context.Param("payId")
		accountId := uint(try(strconv.Atoi(context.Query("accountId"))).(int))
		payState := try(strconv.Atoi(context.Query("state"))).(domain.PayState)
		updatePaymentState(context, payId, accountId, payState)
	})
}

func updatePaymentState(context *gin.Context, payId string, accountId uint, payState domain.PayState) {
	service := payment.WithGinContext(context)
	if payState == domain.Payed {
		response.Op(context, func() {
			if err := service.AccomplishPayment(accountId, payId); err != nil {
				panic(err)
			}
		})
	} else {
		response.Op(context, func() {
			if err := service.CancelPayment(payId); err != nil {
				panic(err)
			}
		})
	}
}
