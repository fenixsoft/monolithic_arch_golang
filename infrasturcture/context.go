package infrasturcture

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	CTXTransaction = "CTX_DB_Transaction"
	CTXUsername    = "CTX_User"
	CTXLogger      = "CTX_Logger"
)

// Transaction 返回由事务中间件自动管理的事务的Session
func Transaction(context *gin.Context) *Database {
	return TxDB(context.Request.Context())
}

// Logger 返回当前上下文中的日志对象
func Logger(context *gin.Context) *logrus.Entry {
	ctx := context.Request.Context()
	return ctx.Value(CTXLogger).(*logrus.Entry)
}

// LoginUser 返回当前登陆的用户名称
func LoginUser(context *gin.Context) string {
	ctx := context.Request.Context()
	return ctx.Value(CTXUsername).(string)
}
