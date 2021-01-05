package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct{}

func (c *ProductController) Register(router gin.IRouter) {
	router.GET("/products", func(context *gin.Context) {
		context.JSON(http.StatusOK, Transactional(context).FindAllProducts())
	})
}
