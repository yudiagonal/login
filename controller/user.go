package controller

import (
	"encoding/json"
	"login/config"
	"login/helper"
	"login/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Login(w http.ResponseWriter, r *http.Request) {
	//ambil inputan user
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	//validasi username ada atau tidak
	var user models.User
	if err := DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "username atau password salah"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
		}

		//validasi password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
			response := map[string]string{"message": "username atau password salah"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}
	}

	//pembuatan token jwt
	expTime := time.Now().Add(time.Minute * 1)
	claim := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "jwt",
			ExpiresAt: &jwt.NumericDate{Time: expTime},
		},
	}

	//alogritma untuk signin
	algoToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	//signed token
	token, err := algoToken.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	//set token ko cooki
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	})

	response := map[string]string{"message": "login success"}
	helper.ResponseJSON(w, http.StatusOK, response)

}

func Register(w http.ResponseWriter, r *http.Request) {
	//ambil inputan dari user
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	//hash password
	hassPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hassPassword)

	//insert ke database
	if err := DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	//hapus token di cookie
	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	response := map[string]string{"message": "logout success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
