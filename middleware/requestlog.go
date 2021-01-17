// 日志中间件
// 提供自动带有上下文访问路径、事务ID等支持
package middleware

import (
	"context"
	ctx2 "github.com/fenixsoft/monolithic_arch_golang/infrasturcture/state"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := logrus.WithFields(logrus.Fields{
			"url": c.Request.URL,
		})
		ctx := c.Request.Context()
		c.Request = c.Request.WithContext(context.WithValue(ctx, ctx2.CTXLogger, logger))
	}
}
