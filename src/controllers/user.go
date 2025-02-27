package controllers

import (
	"crud/src/models"
	"crud/src/services"
	"crud/src/utils"
	"crud/src/validators"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserCreation(c *gin.Context) {
	var userInput validators.UserInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userInput.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: userInput.Password,
		Pincode:  userInput.Pincode,
		City:     userInput.City,
		State:    userInput.State,
		Country:  userInput.Country,
	}

	hashedPassword, err := services.HashPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user.Password = hashedPassword

	token, err := services.GenerateJWT(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate JWT Token",
		})
		return
	}

	result := utils.DB.Create(&user)
	if result.Error != nil {
		if isUniqueConstraintError(result.Error) {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create User"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Created Successfully",
		"user":    user,
		"token":   token,
	})
}

func isUniqueConstraintError(err error) bool {
	if err == nil {
		return false
	}

	// The error message for PostgreSQL unique constraint violations
	// contains the string "duplicate key value violates unique constraint"
	return strings.Contains(err.Error(), "duplicate key value violates unique constraint") ||
		strings.Contains(err.Error(), "UNIQUE constraint failed")
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

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var existingUser models.User

	if err := utils.DB.First(&existingUser, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var userInput validators.UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userInput.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser.Name = userInput.Name
	existingUser.Email = userInput.Email
	existingUser.Pincode = userInput.Pincode
	existingUser.City = userInput.City
	existingUser.State = userInput.State
	existingUser.Country = userInput.Country

	if userInput.Password != "" {
		hashedPassword, err := services.HashPassword(userInput.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to hash password",
			})
			return
		}
		existingUser.Password = hashedPassword
	}

	token, err := services.GenerateJWT(existingUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate JWT Token",
		})
		return
	}

	if err := utils.DB.Save(&existingUser).Error; err != nil {
		if isUniqueConstraintError(err) {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists, please try a different email"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user: " + err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Updated Successfully",
		"user":    existingUser,
		"token":   token,
	})
}

func SoftDeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := utils.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not Found"})
		return
	}

	if err := utils.DB.Delete(&user, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "User Soft Deleted Successfully",
		"deletedUser": user,
	})
}

func HardDeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := utils.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not Found"})
		return
	}

	if err := utils.DB.Unscoped().Delete(&user, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "User Hard Deleted Successfully",
		"deletedUser": user,
	})
}

func UserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := utils.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User fetched Successfully",
		"user":    user,
	})
}

func UserLogin(c *gin.Context) {
	var loginInput validators.LoginInput

	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := loginInput.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := utils.DB.Where("email = ?", loginInput.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email ,account doesn't exists"})
		return
	}

	if err := services.VerifyPassword(user.Password, loginInput.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := services.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT Token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
		"token":   token,
	})
}
