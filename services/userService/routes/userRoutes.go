package routes

import (
	"github.com/gorilla/mux"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userService/handlers"
)

func UserRoutes(router *mux.Router) {
	// CRUD de usuarios
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{user_id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")
	// Autenticación
	router.HandleFunc("/users/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/users/login", handlers.LoginUser).Methods("POST")
}