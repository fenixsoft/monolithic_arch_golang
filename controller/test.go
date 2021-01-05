package controller

import (
	"github.com/fenixsoft/monolithic_arch_golang/domain"
	"github.com/gin-gonic/gin"
)

type TestController struct {
}

func (c *TestController) Register(router gin.IRouter) {
	router.GET("/test", func(c *gin.Context) {
		database := Transactional(c)
		tx := database.Session.Begin()
		database.Session.Find(&domain.Advertisement{})
		tx.Commit()
	})
}
