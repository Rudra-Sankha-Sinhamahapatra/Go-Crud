package main

import (
	"crud/src/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := utils.LoadEnv()

	if err != nil {
		log.Fatal("Error loading env variables")
	}

	utils.InitDB()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
