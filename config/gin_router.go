package config

import (
	controller "github.com/Hasyim-Kai/Go_Gin_Rest_Api/controller/student"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func SetupRouter() *gin.Engine {
	DB := DbInit()

	r := gin.Default()

	r.POST("/student", func(ctx *gin.Context) { controller.PostHandler(ctx, DB) })
	r.GET("/student", func(ctx *gin.Context) { controller.GetAllHandler(ctx, DB) })
	// r.GET("/student/:student_id", func(ctx *gin.Context) { controller.GetHandler(ctx, DB) })
	// r.PUT("/student/:student_id", func(ctx *gin.Context) { controller.PutHandler(ctx, DB) })
	// r.DELETE("/student/:student_id", func(ctx *gin.Context) { controller.DelHandler(ctx, DB) })

	return r
}
