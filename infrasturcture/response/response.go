package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(context *gin.Context, code int, err string) {
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
			ServerError(context, fmt.Sprintf("%v\n", r))
		}
	}()
	fn()
	Success(context)
}

// 搏一搏，行就行，不行就崩
// 用于偷懒，减少if err!=nil的数量
func Try(ret interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return ret
}
