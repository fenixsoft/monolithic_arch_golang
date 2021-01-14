package controller

import (
	"fmt"
	"github.com/fenixsoft/monolithic_arch_golang/domain"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/response"
	"github.com/fenixsoft/monolithic_arch_golang/service"
	"github.com/fenixsoft/monolithic_arch_golang/service/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct{}

func (c *AccountController) Register(router gin.IRouter) {
	// 获取用户信息
	router.GET("/accounts/:username", func(context *gin.Context) {
		username := context.Param("username")
		account, err := domain.GetAccountByName(DB(context), username)
		if err != nil {
			context.String(http.StatusNotFound, "未能找到指定用户")
		} else {
			context.JSON(http.StatusOK, account)
		}
	})

	// 修改用户信息
	router.PUT("/accounts", func(context *gin.Context) {
		if account, err := bindAccount(context); err == nil {
			if validation.RunAccountValidation(context, account, []validation.TypeAccount{validation.Authenticated, validation.NotConflict}) {
				response.Op(context, func() {
					try(domain.UpdateAccount(DB(context), account))
				})
			}
		}
	})

	// 创建新用户
	router.POST("/accounts", func(context *gin.Context) {
		if account, err := bindAccount(context); err == nil {
			if validation.RunAccountValidation(context, account, []validation.TypeAccount{validation.Unique}) {
				response.Op(context, func() {
					account.Password = service.PasswordEncoding(account.Password)
					try(domain.CreateAccount(DB(context), account))
					fmt.Printf("id:%v\n", account.ID)
				})
			}
		}
	})
}

func bindAccount(context *gin.Context) (*domain.Account, error) {
	acc, err := bindModel(context, new(domain.Account))
	return acc.(*domain.Account), err
}
