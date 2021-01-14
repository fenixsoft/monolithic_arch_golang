package controller

import (
	"github.com/fenixsoft/monolithic_arch_golang/domain"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type ProductController struct{}

func (c *ProductController) Register(router gin.IRouter) {
	// 获取仓库中所有的货物信息
	router.GET("/products", func(context *gin.Context) {
		context.JSON(http.StatusOK, try(domain.FindAllProducts(DB(context))))
	})

	// 获取仓库中指定的货物信息
	router.GET("/products/:id", func(context *gin.Context) {
		// 根据ID取货物信息
		context.JSON(http.StatusOK, try(domain.GetProduct(DB(context), toInt(context.Param("id")))))
	})

	// Gin中使用的http router具有冲突限制，不能再注册“/products/stockpile/:id”这样的处理器，因此在代码中处理
	router.GET("/products/:id/:sid", func(context *gin.Context) {
		id := context.Param("id")
		if strings.HasPrefix(id, "stockpile") {
			// 根据ID取库存信息
			context.JSON(http.StatusOK, try(domain.GetStockpile(DB(context), toInt(context.Param("sid")))))
		} else {
			context.AbortWithStatus(http.StatusNotFound)
		}
	})

	// 更新产品信息
	router.PUT("/products", func(context *gin.Context) {
		response.Op(context, func() {
			try(domain.UpdateProduct(DB(context), bindProduct(context)))
		})
	})

	// 创建新的产品
	router.POST("/products", func(context *gin.Context) {
		context.JSON(http.StatusOK, try(domain.CreateProduct(DB(context), bindProduct(context))))
	})

	// 删除产品
	router.DELETE("/products/:id", func(context *gin.Context) {
		response.Op(context, func() {
			try(domain.DeleteProduct(DB(context), toInt(context.Param("id"))))
		})
	})

	// 将指定的产品库存调整为指定数额
	router.PATCH("/products/stockpile/:id", func(context *gin.Context) {
		response.Op(context, func() {
			try(domain.UpdateStockpileAmount(DB(context), toInt(context.Param("id")), toInt(context.Query("amount"))))
		})
	})
}

func toInt(p string) int {
	return try(strconv.Atoi(p)).(int)
}

func bindProduct(context *gin.Context) *domain.Product {
	return try(bindModel(context, new(domain.Product))).(*domain.Product)
}
