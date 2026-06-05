package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId uint64, secret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(secret),
	)
}

func ValidateToken(tokenString string, secret string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
}
