package routes

import (
	"github.com/Hasyim-Kai/Go_Gin_Rest_Api_MySkill/controller"
	"github.com/gin-gonic/gin"
)

func StudentRoutes(router *gin.Engine) {
	route := router.Group("/api/student")
	{
		route.GET("/", controller.PostHandler)
		route.GET("/:id", controller.GetAllHandler)
		route.POST("/", controller.GetHandler)
		route.PUT("/:id", controller.PutHandler)
		route.DELETE("/:id", controller.DelHandler)
	}
}
