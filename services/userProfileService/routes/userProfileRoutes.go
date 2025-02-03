package routes

import (
	"github.com/gorilla/mux"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userProfileService/handlers"
)

func ProfileRoutes(router *mux.Router) {
	router.HandleFunc("/profiles", handlers.GetProfiles).Methods("GET")
	router.HandleFunc("/profiles/{user_id:[0-9]+}", handlers.GetProfile).Methods("GET")
	router.HandleFunc("/profiles", handlers.CreateProfile).Methods("POST")
	router.HandleFunc("/profiles/{user_id:[0-9]+}", handlers.UpdateProfile).Methods("PUT")
	router.HandleFunc("/profiles/{id:[0-9]+}", handlers.DeleteProfile).Methods("DELETE")
}


