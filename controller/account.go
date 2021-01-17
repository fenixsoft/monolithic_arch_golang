package controller

import (
	"github.com/fenixsoft/monolithic_arch_golang/domain"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/response"
	"github.com/fenixsoft/monolithic_arch_golang/service/auth"
	"github.com/fenixsoft/monolithic_arch_golang/service/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct{}

func (c *AccountController) Register(router gin.IRouter) {
	// 获取用户信息
	router.GET("/accounts/:username", func(context *gin.Context) {
		username := context.Param("username")
		account, err := domain.NewAccount(DB(context)).GetByName(username)
		if err != nil {
			context.String(http.StatusNotFound, "未能找到指定用户")
		} else {
			context.JSON(http.StatusOK, account)
		}
	})

	// 修改用户信息
	router.PUT("/accounts", func(context *gin.Context) {
		account := bindAccount(context)
		if validation.RunAccountValidation(context, account, []validation.TypeAccount{validation.Authenticated, validation.NotConflict}) {
			response.Op(context, func() {
				try(account.Update())
			})
		}
	})

	// 创建新用户
	router.POST("/accounts", func(context *gin.Context) {
		account := bindAccount(context)
		if validation.RunAccountValidation(context, account, []validation.TypeAccount{validation.Unique}) {
			response.Op(context, func() {
				account.Password = auth.PasswordEncoding(account.Password)
				try(account.Create())
			})
		}
	})
}

func bindAccount(context *gin.Context) *domain.Account {
	return try(bindModel(context, domain.NewAccount(DB(context)))).(*domain.Account)
}
