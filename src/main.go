package main

import (
	"crud/src/controllers"
	"crud/src/utils"
	"log"

	_ "crud/src/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	r.POST("/create-user", controllers.UserCreation)
	r.GET("/getById/:id", controllers.UserById)
	r.GET("/all-users", controllers.AllUser)
	r.PUT("/update-user/:id", controllers.UpdateUser)
	r.DELETE("/delete-user/:id", controllers.DeleteUser)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
