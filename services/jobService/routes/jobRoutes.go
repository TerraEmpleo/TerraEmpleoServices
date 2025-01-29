package routes

import (
	"github.com/gorilla/mux"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/jobService/handlers"
)

func JobRoutes(router *mux.Router) {
	router.HandleFunc("/jobs", handlers.GetJobs).Methods("GET")
	router.HandleFunc("/jobs/{id}", handlers.GetJob).Methods("GET")
	router.HandleFunc("/jobs", handlers.CreateJob).Methods("POST")
	router.HandleFunc("/jobs/{id}", handlers.UpdateJob).Methods("PUT")
	router.HandleFunc("/jobs/{id}", handlers.DeleteJob).Methods("DELETE")
}