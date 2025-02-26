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

	// Gin mode based on environment
	if utils.AppConfig.ENV == "production" {
		log.Println("Running in production mode")
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.Println("Running in development mode")
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/create-user", controllers.UserCreation)
	r.POST("/login", controllers.UserLogin)
	r.GET("/getById/:id", controllers.UserById)
	r.GET("/all-users", controllers.AllUser)
	r.PUT("/update-user/:id", controllers.UpdateUser)
	r.DELETE("/soft/delete-user/:id", controllers.SoftDeleteUser)
	r.DELETE("/hard/delete-user/:id", controllers.HardDeleteUser)

	// Only register Swagger in development mode
	if utils.AppConfig.ENV == "development" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		log.Println("Swagger UI enabled at /swagger/index.html")
	} else {
		log.Println("Swagger UI disabled in production mode")
	}

	r.Run()
}
