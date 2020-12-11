package main

import (
	"github.com/gin-gonic/gin"
	"tanglei_1211/messageboard/lv1/listen"
)

func main()  {
	router := gin.Default()
	router.POST("/register",listen.Register)
	_ = router.Run(":8080")
}
