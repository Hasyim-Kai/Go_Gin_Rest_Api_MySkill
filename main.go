package main

import "github.com/Hasyim-Kai/Go_Gin_Rest_Api/config"

func main() {
	// practice.RunArray()
	// practice.RunDeferPanicError()
	// practice.RunJson()
	// practice.RunInterface()
	// practice_goroutine.RunGoroutines()

	r := config.SetupRouter()
	r.Run(":8080")
}
