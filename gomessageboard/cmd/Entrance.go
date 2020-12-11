package cmd

import (
	"github.com/gin-gonic/gin"
	"gomessageboard/controller"
)

func Entrance() {
	router := gin.Default()
	router.POST("/register",controller.Register)
	router.GET("/login",controller.Login)
	router.POST("/postMessage",controller.SendMessage)
	router.GET("/viewMessage",controller.ViewMessage)
	router.Run(":8080")
}
