// 静态资源控制器，负责注册HTML、CSS、图片等静态资源
package controller

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StaticController struct{}

func (c *StaticController) Register(router gin.IRouter) {
	// 注册静态资源
	// 出于httprouter的限制，分别注册根目录和/static目录的静态资源
	router.StaticFS("/static", rice.MustFindBox("../resource/static/static").HTTPBox())
	router.GET("/", func(c *gin.Context) {
		index, _ := rice.MustFindBox("../resource/static").Bytes("index.html")
		c.Data(http.StatusOK, "text/html", index)
	})
	router.GET("/favicon.ico", func(c *gin.Context) {
		index, _ := rice.MustFindBox("../resource/static").Bytes("favicon.ico")
		c.Data(http.StatusOK, "image/x-icon", index)
	})
}
