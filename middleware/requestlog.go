package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	LoggerContext = "LOGGER_CTX"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := logrus.WithFields(logrus.Fields{
			"url": c.Request.URL,
		})
		ctx := c.Request.Context()
		c.Request = c.Request.WithContext(context.WithValue(ctx, LoggerContext, logger))
	}
}

func Logger(context *gin.Context) *logrus.Entry {
	ctx := context.Request.Context()
	return ctx.Value(LoggerContext).(*logrus.Entry)
}
