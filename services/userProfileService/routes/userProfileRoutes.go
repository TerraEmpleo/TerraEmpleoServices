package routes

import (
	"github.com/gorilla/mux"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userProfileService/handlers"
)

func ProfileRoutes(router *mux.Router) {
	router.HandleFunc("/profiles", handlers.GetProfiles).Methods("GET")
	router.HandleFunc("/profiles/{id}", handlers.GetProfile).Methods("GET")
	router.HandleFunc("/profiles", handlers.CreateProfile).Methods("POST")
	router.HandleFunc("/profiles/{id}", handlers.UpdateProfile).Methods("PUT")
	router.HandleFunc("/profiles/{id}", handlers.DeleteProfile).Methods("DELETE")
}
