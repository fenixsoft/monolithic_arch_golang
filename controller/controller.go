// 控制器公共方法
package controller

import (
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/db"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/response"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/state"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/util"
	"github.com/fenixsoft/monolithic_arch_golang/middleware"
	"github.com/gin-gonic/gin"
)

// 系统中所有控制器的列表，目前以列表形式，以后多起来可能需要个配置或者DI框架
var (
	static StaticController
	oauth  OAuth2Controller
	rests  = [...]IController{
		new(AdvertisementController),
		new(ProductController),
		new(AccountController),
		new(SettlementController),
		new(PaymentController),
	}
)

type IController interface {
	Register(gin.IRouter)
}

// Register 统一将系统中所有的控制器注册到路由中
func Register(router *gin.Engine) {
	static.Register(router)
	oauth.Register(router)

	restGroup := router.Group("/restful")
	// restful路径下中间件的支持，目前支持事务、请求日志、鉴权
	// 由于其他中间件中会使用到日志，所以日志应该放第一个
	restGroup.Use(middleware.RequestLoggerMiddleware())
	restGroup.Use(middleware.TransitionalMiddleware())
	restGroup.Use(middleware.SecurityMiddleware())
	for _, v := range rests {
		v.Register(restGroup)
	}
}

func bindModel(context *gin.Context, model interface{}) (interface{}, error) {
	if err := context.ShouldBindJSON(model); err != nil {
		response.ClientError(context, err.Error())
		return nil, err
	}
	return model, nil
}

// 起个别名
var (
	// DB  = state.Database
	try = util.Try
)

func DB(context *gin.Context) *db.Database {
	return state.WithGinContext(context).Database()
}
