package controller

import (
	"github.com/fenixsoft/monolithic_arch_golang/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdvertisementController struct{}

func (c *AdvertisementController) Register(router gin.IRouter) {
	router.GET("/advertisements", func(context *gin.Context) {
		context.JSON(http.StatusOK, middleware.Transactional(context).FindAllAdvertisements())
	})
}
