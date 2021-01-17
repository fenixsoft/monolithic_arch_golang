package util

import "github.com/sirupsen/logrus"

// 搏一搏，行就行，不行就崩
// 用于偷懒，减少if err!=nil的数量
// 注意只在中间件、页面控制器等外层必定有recover兜底的场合中使用
func Try(ret interface{}, err error) interface{} {
	if err != nil {
		logrus.StandardLogger().Warn(err.Error())
		panic(err)
	}
	return ret
}
