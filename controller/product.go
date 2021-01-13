package controller

import (
	"github.com/fenixsoft/monolithic_arch_golang/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductController struct{}

func (c *ProductController) Register(router gin.IRouter) {
	// 获取所有产品
	router.GET("/products", func(context *gin.Context) {
		context.JSON(http.StatusOK, domain.FindAllProducts(DB(context)))
	})

	router.GET("/products/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		context.JSON(http.StatusOK, domain.GetProduct(DB(context), id))
	})
}
