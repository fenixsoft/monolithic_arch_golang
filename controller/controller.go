// 控制器公共方法
package controller

import (
	"github.com/fenixsoft/monolithic_arch_golang/middleware"
	"github.com/gin-gonic/gin"
)

// 系统中所有控制器的列表
var (
	static StaticController
	rests  = [...]IController{
		new(AdvertisementController),
		new(ProductController),
		new(TestController),
	}
)

type IController interface {
	Register(router gin.IRouter)
}

// Register 统一将系统中所有的控制器注册到路由中
func Register(router *gin.Engine) {
	static.Register(router)

	restGroup := router.Group("/restful")

	// restful路径下中间件的支持，目前支持事务和请求日志
	// 其他中间件中会使用到日志，所以日志应该放第一个
	restGroup.Use(middleware.RequestLoggerMiddleware())
	restGroup.Use(middleware.TransitionalMiddleware())

	for _, v := range rests {
		v.Register(restGroup)
	}
}
