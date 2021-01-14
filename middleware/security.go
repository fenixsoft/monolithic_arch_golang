package middleware

import (
	"context"
	ctx2 "github.com/fenixsoft/monolithic_arch_golang/infrasturcture/ctx"
	"github.com/fenixsoft/monolithic_arch_golang/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SecurityMiddleware() gin.HandlerFunc {
	return func(gc *gin.Context) {
		auth := gc.Request.Header.Get("Authorization")
		// 判断是否以“bearer ”开头，部分服务是允许没有令牌访问的，只是拿不到上下文中的username了
		if strings.HasPrefix(auth, "bearer ") {
			claims, err := service.ValidatingJWTAccessToken(auth)
			if err != nil {
				_ = gc.AbortWithError(http.StatusUnauthorized, err)
			}
			// 验证令牌有效，从令牌中取出username（其他像Scope这些，在Golang版本中就不使用了）
			ctx := gc.Request.Context()
			gc.Request = gc.Request.WithContext(context.WithValue(ctx, ctx2.CTXUsername, claims["username"]))
		}
	}
}
