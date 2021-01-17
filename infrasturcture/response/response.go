package response

import (
	"errors"
	"fmt"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/db"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/state"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(context *gin.Context, code int, err string) {
	d := state.WithGinContext(context).Database()
	d.Error = errors.New(err)
	d.State = db.TransactionStateRollback
	context.AbortWithStatusJSON(code, gin.H{"code": 1, "message": err})
}

func ClientError(context *gin.Context, err string) {
	Error(context, http.StatusBadRequest, err)
}

func ServerError(context *gin.Context, err string) {
	Error(context, http.StatusInternalServerError, err)
}

func Success(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作已成功"})
}

// 将函数出错截获，以异常信息返回给前端
func Op(context *gin.Context, fn func()) {
	defer func() {
		if r := recover(); r != nil {
			state.WithGinContext(context).Logger().Errorf("%v\n", r)
			ServerError(context, fmt.Sprintf("%v\n", r))
		}
	}()
	fn()
	state.WithGinContext(context).Database().State = db.TransactionStateCommit
	Success(context)
}
