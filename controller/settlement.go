// 结算清单相关的资源
package controller

import (
	"github.com/fenixsoft/monolithic_arch_golang/domain"
	"github.com/fenixsoft/monolithic_arch_golang/service/payment"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SettlementController struct{}

func (c *SettlementController) Register(router gin.IRouter) {
	router.POST("/settlements", func(context *gin.Context) {
		p := payment.New().ExecuteBySettlement(bindSettlement(context))
		context.JSON(http.StatusOK, p)
	})
}

func bindSettlement(context *gin.Context) *domain.Settlement {
	return try(bindModel(context, &domain.Settlement{ProductMap: make(map[uint]*domain.Product)})).(*domain.Settlement)

}
