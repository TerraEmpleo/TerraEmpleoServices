package routes

import (
	"github.com/gorilla/mux"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/applicationService/handlers"
)

func ApplicationRoutes(router *mux.Router) {
	router.HandleFunc("/applications/apply", handlers.ApplyForJob).Methods("POST")
	router.HandleFunc("/applications", handlers.GetApplications).Methods("GET")
}
