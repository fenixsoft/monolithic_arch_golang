// 事务中间件
// 请求经过业务方法（即/restful的服务）时，会在事务中间件中开启事务，并根据是否抛出了panic来自动提交或回滚事务
// 仅提供最基础最简单的事务自动回滚与提交、回滚。没有去理会管理器、只读事务、嵌套、传播级别、自定义出错类型等功能。
// 这种类似于Java Hibernate中OpenSessionInView的设计并不见得是最佳的，按照Golang提倡的风格，编程式事务才是首选。

package middleware

import (
	"context"
	ctx2 "github.com/fenixsoft/monolithic_arch_golang/infrasturcture/ctx"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/db"
	"github.com/gin-gonic/gin"
)

func TransitionalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		tx := db.DB.Session.WithContext(ctx).Begin()
		c.Request = c.Request.WithContext(context.WithValue(ctx, db.CTXTransaction, tx))
		logger := ctx2.Logger(c)
		logger.WithField("Tx", db.TxID(tx.Statement)).Trace("开启中间件事务")
		defer func() {
			if r := recover(); r != nil {
				logger.WithField("Tx", db.TxID(tx.Statement)).Errorf("回滚中间件事务，异常原因：%v\n", r)
				tx.Rollback()
				// 不在事务中间件中处理恐慌，回滚后继续抛出恐慌，在后续的Recovery中间件中统一解决
				panic(r)
			}
		}()
		c.Next()
		logger.WithField("Tx", db.TxID(tx.Statement)).Trace("提交中间件事务")
		tx.Commit()
	}
}
