// 控制器公共方法
package controller

import (
	"context"
	"fmt"
	"github.com/fenixsoft/monolithic_arch_golang/domain"
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
	restGroup.Use(TransitionalMiddleware())
	for _, v := range rests {
		v.Register(restGroup)
	}
}

// Transactional 返回由事务中间件自动管理的事务的Session
func Transactional(context *gin.Context) *domain.Database {
	return domain.DB.Context(context.Request.Context())
}

// TransitionalMiddleware 事务中间件
// 请求经过业务方法（即/restful的服务）时，会在事务中间件中开启事务，并根据是否抛出了panic来自动提交或回滚事务
// 仅提供最基础最简单的事务自动回滚与提交、回滚。没有去理会管理器、只读事务、嵌套、传播级别、自定义出错类型等功能。
func TransitionalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		tx := domain.DB.Session.WithContext(ctx).Begin()
		c.Request = c.Request.WithContext(context.WithValue(ctx, domain.TransactionContext, tx))

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		c.Next()
		fmt.Printf("%v\n", &tx.Statement.ConnPool)
		tx.Commit()
	}
}
