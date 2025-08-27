package controllers

import (
	"auth/src/views"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	views.JSON(w, http.StatusOK, nil)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	views.JSON(w, http.StatusOK, nil)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	views.JSON(w, http.StatusOK, nil)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	views.JSON(w, http.StatusOK, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	views.JSON(w, http.StatusOK, nil)
}
