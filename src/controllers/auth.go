package controllers

import (
	"encoding/json"
	"net/http"

	database "auth/config"
	"auth/src/auth"
	"auth/src/models"
	"auth/src/views"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	type FormLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var form FormLogin
	var user models.User

	// Decode body
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		views.Message(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	// Find user by email
	if database.Conn.Where("email = ?", form.Email).First(&user).Error != nil {
		views.Message(w, http.StatusUnauthorized, "invalid email or password")
		return
	}

	// Check password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)) != nil {
		views.Message(w, http.StatusUnauthorized, "invalid email or password")
		return
	}

	// Generate token
	token, err := auth.GenToken(user)
	if err != nil {
		views.Message(w, http.StatusInternalServerError, "could not generate token")
		return
	}

	// Send response
	views.JSON(w, http.StatusOK, map[string]interface{}{"token": token, "user": user.ResponseFormat()})
}