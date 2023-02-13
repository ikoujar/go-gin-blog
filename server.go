package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-course/common/config"
	"go-gin-course/common/middleware"
	"go-gin-course/pkg/comment"
	"go-gin-course/pkg/post"
	"log"
)

func main() {

	config.Init()

	server := gin.Default()

	server.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// Server middlewares.
	server.Use(middleware.LoggerMiddleware)
	server.Use(middleware.ErrorHandler())

	// Routes initialization.
	initRoutes(server)

	// Start the server.
	err := server.Run(":8000")
	if err != nil {
		log.Fatal("An error occurred while starting the application")
		return
	}
}

func initRoutes(route *gin.Engine) {
	post.Init(route)
	comment.Init(route)
}
