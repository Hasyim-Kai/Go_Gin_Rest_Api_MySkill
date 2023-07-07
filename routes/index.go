package routes

import "github.com/gin-gonic/gin"

func IndexRoutes(router *gin.Engine) {
	StudentRoutes(router)
}
