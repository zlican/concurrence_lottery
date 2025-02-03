package router

import (
	"lottery/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/home", controller.GetHome)
		api.GET("/prize/all", controller.GetAllPrize)
		api.GET("/prize/lottery", controller.Lottery)
	}
}
