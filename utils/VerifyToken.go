package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	UID int    `json:"uid"`
	Exp int64  `json:"exp"`
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Aud string `json:"aud"`
	Iat int64  `json:"iat"`
	Nbf int64  `json:"nbf"`
	jwt.StandardClaims
}

func VerifyToken(authHeader string) (*JwtClaims, error) {
	prefix := "Bearer "
	tokenString := strings.TrimPrefix(authHeader, prefix)
	jwtKey := GetJwtSecret()

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("authentication header not present or malformed")
	}

	claims, ok := token.Claims.(*JwtClaims)

	if ok {
		if claims.Exp != 0 && claims.Exp < time.Now().Unix() {
			return nil, fmt.Errorf("token is expired")
		} else {
			return claims, nil
		}
	}

	return claims, nil
}
