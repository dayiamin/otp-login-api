package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))


// generate jwt token for user and put userID and user Phone in its header with 24h expiry
func GenerateJWT(phone string) (string, error) {
	claims := jwt.MapClaims{
		"phone_number": phone,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 24h expiry
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JwtSecret)
}
