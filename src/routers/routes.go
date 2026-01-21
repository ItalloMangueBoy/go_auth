package routes

import (
	"auth/src/controllers"

	"github.com/gorilla/mux"
)

// SetupRoutes initializes the application's routes.
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.ListUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/login", controllers.Login).Methods("POST")

	return router
}
