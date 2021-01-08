package controller

import (
	"github.com/fenixsoft/monolithic_arch_golang/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct{}

func (c *ProductController) Register(router gin.IRouter) {
	router.GET("/products", func(context *gin.Context) {
		context.JSON(http.StatusOK, middleware.Transactional(context).FindAllProducts())
	})
}
