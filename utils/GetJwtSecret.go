package utils

import "os"

type JwtSecret struct {
	Secret []byte `json:"secret"`
}

func GetJwtSecret() []byte {
	jwt := &JwtSecret{}
	jwt.Secret = []byte(os.Getenv("JWT_SIGN_KEY"))

	return jwt.Secret
}
