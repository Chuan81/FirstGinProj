package main

import (
	"exchangeapp/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	type Info struct {
		Message string
	}

	InfoTest := Info{
		Message: "pong",
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, InfoTest)
	})
	r.Run(config.AppConfig.App.Port)
}
