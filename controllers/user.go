package controllers

import (
	"artiz-pogo-api/models"
	"artiz-pogo-api/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

type JwtSecret struct {
	Secret []byte `json:"secret"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please provide the correct input!!"))
		return
	}

	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)

	if errHash != nil {
		http.Error(w, "Could not generate password hash", http.StatusBadRequest)
		return
	}

	var existingUser models.User
	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	} else {
		models.DB.Create(&user)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please provide the correct input!!"))
		return
	}

	var existingUser models.User
	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID == 0 {
		http.Error(w, "User does not exist", http.StatusBadRequest)
		return
	}

	errHash := utils.CompareHashPassword(existingUser.Password, user.Password)
	if !errHash {
		http.Error(w, "Incorrect password please check again", http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "localhost:8080",
		"sub": "localhost",
		"aud": "*",
		"uid": existingUser.ID,
		"exp": expirationTime,
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
	})

	jwtKey := GetJwtSecret()
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Claims could not be set", http.StatusBadRequest)
		return
	}

	jsonObject := map[string]interface{}{
		"token": tokenString,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsonObject)
}

func CurrentUser(w http.ResponseWriter, r *http.Request) {
	token, err := VerifyToken(r.Header.Get("Authorization"))
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	jsonObject := map[string]interface{}{
		"userId": token.UID,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsonObject)
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

func GetJwtSecret() []byte {
	jwt := &JwtSecret{}
	jwt.Secret = []byte(os.Getenv("JWT_SIGN_KEY"))

	return jwt.Secret
}
