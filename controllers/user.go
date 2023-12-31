package controllers

import (
	"artiz-pogo-api/models"
	"artiz-pogo-api/utils"
	"encoding/json"
	"fmt"
	"net/http"

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

	tokenString, err := utils.GenerateToken(existingUser.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonObject := map[string]interface{}{
		"token": tokenString,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsonObject)
}

func CurrentUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("userId")

	fmt.Println("=====userId====")
	fmt.Println(userId)

	jsonObject := map[string]interface{}{
		"userId": userId,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsonObject)
}
