package validation

import (
	"github.com/fenixsoft/monolithic_arch_golang/domain"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/response"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/state"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/util"
	"github.com/gin-gonic/gin"
)

type TypeAccount int

const (
	NotConflict TypeAccount = iota
	Authenticated
	Unique
)

func RunAccountValidation(context *gin.Context, user *domain.Account, types []TypeAccount) bool {
	var r = true
	for i := 0; i < len(types) && r; i++ {
		switch types[i] {
		case NotConflict:
			r = r && NotConflictAccount(context, user)
		case Authenticated:
			r = r && AuthenticatedAccount(context, user)
		case Unique:
			r = r && UniqueAccount(context, user)
		}
	}
	return r
}

// 表示一个用户的信息是无冲突的
// “无冲突”是指该用户的敏感信息与其他用户不重合，譬如将一个注册用户的邮箱，修改成与另外一个已存在的注册用户一致的值，这便是冲突
func NotConflictAccount(context *gin.Context, user *domain.Account) bool {
	accounts := util.Try(domain.NewAccount(state.WithGinContext(context).Database()).FindByNameOrEmailOrTelephone(user.Username, user.Email, user.Telephone)).([]domain.Account)
	size := len(accounts)
	if !(size == 0 || (size == 1 && accounts[0].ID == user.ID)) {
		response.ClientError(context, "用户个人资料冲突")
		return false
	}
	return true
}

// 代表用户必须与当前登陆的用户一致
func AuthenticatedAccount(context *gin.Context, user *domain.Account) bool {
	if !(state.WithGinContext(context).LoginUser() == user.Username) {
		response.ClientError(context, "用户未登陆")
		return false
	}
	return true
}

// 表示一个用户是唯一的
// 唯一不仅仅是用户名，还要求手机、邮箱均不允许重复
func UniqueAccount(context *gin.Context, user *domain.Account) bool {
	accounts := util.Try(domain.NewAccount(state.WithGinContext(context).Database()).FindByNameOrEmailOrTelephone(user.Username, user.Email, user.Telephone)).([]domain.Account)
	if len(accounts) != 0 {
		response.ClientError(context, "用户名、邮件、电话改成与现有存在重复")
		return false
	}
	return true
}
