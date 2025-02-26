package services

import (
	"crud/src/models"
	"crud/src/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(utils.AppConfig.JWT_SECRET)

func GenerateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(jwtSecret)
}
