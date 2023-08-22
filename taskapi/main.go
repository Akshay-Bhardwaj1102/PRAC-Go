package main

import (
	"github.com/gin-gonic/gin"
	"taskapi.com/config"
	"taskapi.com/controller"
)

func main() {

	r := gin.Default()

	// loading env file
	config.LoadEnv()

	// connecting to database
	config.ConnectDatabase()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong Akshay...",
	// 	})
	// })

	r.GET("/All", controller.FindAll)
	r.POST("/Add", controller.CreateLeave)
	r.GET("/All/:name", controller.FindAllByName)
	r.GET("/Find/:name", controller.FindLeave)
	r.DELETE("/Del/:name", controller.DeleteLeave)

	defer r.Run()
}
