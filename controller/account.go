package controller

import (
	"github.com/fenixsoft/monolithic_arch_golang/domain"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture"
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
		account, num := domain.GetAccountByName(DB(context), username)
		if num == 0 {
			context.String(http.StatusNotFound, "未能找到指定用户")
		} else {
			context.JSON(http.StatusOK, account)
		}
	})

	// 修改用户信息
	router.PUT("/accounts", func(context *gin.Context) {
		if account, err := bindAccount(context); err == nil {
			if validation.RunAccountValidation(context, account, []validation.TypeAccount{validation.Authenticated, validation.NotConflict}) {
				infrasturcture.Op(context, func() {
					domain.UpdateAccount(DB(context), account)
				})
			}
		}
	})

	// 创建新用户
	router.POST("/accounts", func(context *gin.Context) {
		if account, err := bindAccount(context); err == nil {
			if validation.RunAccountValidation(context, account, []validation.TypeAccount{validation.Unique}) {
				infrasturcture.Op(context, func() {
					account.Password = service.PasswordEncoding(account.Password)
					domain.CreateAccount(DB(context), account)
				})
			}
		}
	})
}

func bindAccount(context *gin.Context) (*domain.Account, error) {
	var account domain.Account
	if err := context.ShouldBindJSON(&account); err != nil {
		infrasturcture.ClientError(context, err.Error())
		return nil, err
	}
	return &account, nil
}
