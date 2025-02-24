package controllers

import (
	"crud/src/models"
	"crud/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserCreation(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := utils.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Created Successfully",
		"user":    user,
	})
}

func AllUser(c *gin.Context) {
	var users []models.User

	result := utils.DB.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch all users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully fetched All Users",
		"users":   users,
	})
}
