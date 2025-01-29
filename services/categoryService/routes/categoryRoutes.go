package routes

import (
	"github.com/gorilla/mux"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/categoryService/handlers"
)

func CategoryRoutes(router *mux.Router) {
	router.HandleFunc("/categories", handlers.GetCategories).Methods("GET")
	router.HandleFunc("/categories/{id}", handlers.GetCategory).Methods("GET")
	router.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/{id}", handlers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id}", handlers.DeleteCategory).Methods("DELETE")
}
