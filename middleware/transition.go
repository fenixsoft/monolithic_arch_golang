// 事务中间件
// 请求经过业务方法（即/restful的服务）时，会在事务中间件中开启事务，并根据是否抛出了panic来自动提交或回滚事务
// 仅提供最基础最简单的事务自动回滚与提交、回滚。没有去理会管理器、只读事务、嵌套、传播级别、自定义出错类型等功能。
// 这种类似于Java Hibernate中OpenSessionInView的设计并不见得是最佳的，按照Golang提倡的风格，编程式事务才是首选。

package middleware

import (
	"context"
	"github.com/fenixsoft/monolithic_arch_golang/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const TransactionContext = "DB_CTX"

func TransitionalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		tx := domain.DB.Session.WithContext(ctx).Begin()
		c.Request = c.Request.WithContext(context.WithValue(ctx, TransactionContext, tx))
		logger := Logger(c)
		logger.Debugf("开启中间件事务：%v\n", domain.ConnTrace(tx.Statement))
		defer func() {
			if r := recover(); r != nil {
				logger.Debugf("回滚中间件事务：%v\n", domain.ConnTrace(tx.Statement))
				tx.Rollback()
			}
		}()
		c.Next()
		logger.Debugf("提交中间件事务：%v\n", domain.ConnTrace(tx.Statement))
		tx.Commit()
	}
}

// Transactional 返回由事务中间件自动管理的事务的Session
func Transactional(context *gin.Context) *domain.Database {
	ctx := context.Request.Context()
	session := ctx.Value(TransactionContext).(*gorm.DB)
	ctxDB := new(domain.Database)
	*ctxDB = domain.DB
	ctxDB.Session = session
	return ctxDB
}
