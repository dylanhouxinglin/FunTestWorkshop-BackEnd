package controller

import (
	"FunTestWorkshop/service"
	"github.com/gin-gonic/gin"
	"log"
)

var engine *gin.Engine

func init() {
	engine = gin.Default()
}

func RegControllers() {
	defer func() {
		err := engine.Run(":8000")
		if err != nil {
			log.Fatalln("Server Start Error")
			return
		}
	}()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	engine.POST("/updateRank", service.UpdateRank)
}
