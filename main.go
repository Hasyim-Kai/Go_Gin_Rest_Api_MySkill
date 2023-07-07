package main

import (
	"github.com/Hasyim-Kai/Go_Gin_Rest_Api_MySkill/config"
	"github.com/Hasyim-Kai/Go_Gin_Rest_Api_MySkill/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// practice.RunArray()
	// practice.RunDeferPanicError()
	// practice.RunJson()
	// practice.RunInterface()
	// practice_goroutine.RunGoroutines()

	r := gin.Default()
	config.ConnectDatabase()
	routes.IndexRoutes(r)
	r.Run(":8080")
}
