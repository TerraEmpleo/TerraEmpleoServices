package routes

import (
	"github.com/gorilla/mux"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/jobService/handlers"
)

func JobRoutes(router *mux.Router) {
	router.HandleFunc("/jobs", handlers.GetJobs).Methods("GET")
	router.HandleFunc("/jobs/{id}", handlers.GetJob).Methods("GET")
	router.HandleFunc("/jobs/{user_id:[0-9]+}", handlers.CreateJob).Methods("POST")
	router.HandleFunc("/jobs/{id}", handlers.UpdateJob).Methods("PUT")
	router.HandleFunc("/jobs/{id}", handlers.DeleteJob).Methods("DELETE")
	router.HandleFunc("/jobs/search", handlers.SearchJobs).Methods("GET")
	router.HandleFunc("/jobs/user/{user_id}", handlers.GetJobsByUser).Methods("GET")
	router.HandleFunc("/jobs/recommend/location/{user_id}", handlers.RecommendJobsByLocationForFarmer).Methods("GET")
	router.HandleFunc("/jobs/recommend/farmers/{job_id}", handlers.RecommendFarmersForEmployer).Methods("GET")
}