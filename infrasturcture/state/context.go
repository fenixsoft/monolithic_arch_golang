package state

import (
	"context"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/db"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	CTXUsername    = "CTX_User"
	CTXLogger      = "CTX_Logger"
	CTXTransaction = "CTX_Transaction"
)

type STDCtx struct {
	context context.Context
	db      *db.Database
	logger  *logrus.Entry
	user    string
}

// 与上下文状态无关应用，取全局的数据库连接和日志
func New() *STDCtx {
	return &STDCtx{
		context: nil,
		db:      db.NewDB(),
		logger:  logrus.StandardLogger().WithFields(logrus.Fields{}),
		user:    "",
	}
}

func WithContext(ctx context.Context) *STDCtx {
	return &STDCtx{
		context: ctx,
		logger:  ctx.Value(CTXLogger).(*logrus.Entry),
		db:      getTxDB(ctx),
		user:    loginUser(ctx),
	}
}

func WithGinContext(ctx *gin.Context) *STDCtx {
	return WithContext(ctx.Request.Context())
}

// 返回由事务中间件自动管理的事务的Session
func (ctx *STDCtx) Database() *db.Database {
	return ctx.db
}

// 返回当前上下文中的日志对象
func (ctx *STDCtx) Logger() *logrus.Entry {
	return ctx.logger
}

// 返回当前登陆的用户名称
func (ctx *STDCtx) LoginUser() string {
	return ctx.user
}

// 返回当前上下文
func (ctx *STDCtx) Context() context.Context {
	return ctx.context
}

func loginUser(ctx context.Context) string {
	if r, ok := ctx.Value(CTXUsername).(string); ok {
		return r
	} else {
		return ""
	}
}

func getTxDB(ctx context.Context) *db.Database {
	// 如果上下文中有数据库会话，就使用上下文的，否则就使用全局的
	if ctxDB, ok := ctx.Value(db.CTXDatabase).(*db.Database); ok {
		return ctxDB
	} else {
		return db.NewDB()
	}
}
