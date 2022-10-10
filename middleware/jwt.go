package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// fungsi jwt create token sama extract token

//Create token with jwt

func CreateTokenJwt(id uint, admin bool, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"id":      id,
		"isAdmin": true,
		"exp":     time.Now().Add(12 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
