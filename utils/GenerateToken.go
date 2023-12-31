package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userId uint) (*string, error) {
	expirationTime := time.Now().Add(5 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "localhost:8080",
		"sub": "localhost",
		"aud": "*",
		"uid": userId,
		"exp": expirationTime,
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
	})

	jwtKey := GetJwtSecret()
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, fmt.Errorf("claims could not be set")
	}

	return &tokenString, nil
}
