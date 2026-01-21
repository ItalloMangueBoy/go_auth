package controllers

import (
	database "auth/config"
	"auth/src/models"
	"auth/src/views"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	database.Conn.Omit("password")

	// Decode body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		views.Message(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	// Validate data
	if err := user.Format().Valid(); err != nil {
		views.JSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		views.Message(w, http.StatusInternalServerError, "could not hash password")
		return
	}

	user.Password = string(hash)

	// Save in database
	if err := database.Conn.Create(&user).Error; err != nil {
		views.Message(w, http.StatusInternalServerError, "could not create user")
		return
	}

	// Send response
	views.JSON(w, http.StatusCreated, user.ResponseFormat())
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	// define search query
	search := strings.TrimSpace(r.URL.Query().Get("search"))
	search = "%" + strings.ToLower(search) + "%"

	query := database.Conn.Where("LOWER(name) LIKE ?", search)

	// Find in database
	if err := query.Find(&users).Error; err != nil {
		views.Message(w, http.StatusInternalServerError, "could not retrieve users")
		return
	}

	// Send response
	if len(users) == 0 {
		views.JSON(w, http.StatusNotFound, "no users found")
		return
	}

	views.JSON(w, http.StatusOK, models.UsersToResponseFormat(users))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	vars := mux.Vars(r)

	// Get id from request
	id, error := strconv.ParseUint(vars["id"], 10, 64)
	if error != nil {
		views.Message(w, http.StatusBadRequest, "invalid user id")
		return
	}

	// Find user in database
	if err := database.Conn.First(&user, id).Error; err != nil {
		views.Message(w, http.StatusNotFound, "user not found")
		return
	}

	// Send response
	views.JSON(w, http.StatusOK, user.ResponseFormat())
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get user from URL
	var user models.User

	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		views.Message(w, http.StatusBadRequest, "invalid user id")
		return
	}

	err = database.Conn.First(&user, id).Error
	if err != nil {
		views.Message(w, http.StatusNotFound, "user not found")
		return
	}

	// Get updates from request body
	var updates models.User

	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		views.Message(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	if err := updates.Format().Valid(); err != nil {
		views.JSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(updates.Password), 12)
	if err != nil {
		views.Message(w, http.StatusInternalServerError, "could not hash password")
		return
	}

	updates.Password = string(hash)


	// Update user in database
	if err := database.Conn.Model(&user).Updates(updates).Error; err != nil {
		views.Message(w, http.StatusInternalServerError, "could not update user")
		return
	}

	err = database.Conn.First(&user, id).Error
	if err != nil {
		views.Message(w, http.StatusInternalServerError, "could not reload user")
		return
	}

	// Send response
	views.JSON(w, http.StatusOK, user.ResponseFormat())
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get id from URL
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		views.Message(w, http.StatusBadRequest, "invalid user id")
		return
	}

	// Deletes user
	res := database.Conn.Delete(&models.User{}, id)
	if res.Error != nil {
		views.Message(w, http.StatusInternalServerError, "could not delete user")
		return
	}

	if res.RowsAffected == 0 {
		views.Message(w, http.StatusNotFound, "user not found")
		return
	}

	// Send response
	w.WriteHeader(http.StatusNoContent)
}
