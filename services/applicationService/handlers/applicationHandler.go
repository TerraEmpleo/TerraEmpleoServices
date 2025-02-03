package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/applicationService/database"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/applicationService/models"
	"github.com/gorilla/mux"
)

// Aplicar a un trabajo
func ApplyForJob(w http.ResponseWriter, r *http.Request) {
	var userApplication models.UserApplication

	// Decodificar JSON del request
	if err := json.NewDecoder(r.Body).Decode(&userApplication); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insertar el registro en la tabla intermedia
	if err := database.DB.Create(&userApplication).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userApplication)
}

// Obtener todas las aplicaciones
func GetApplications(w http.ResponseWriter, r *http.Request) {
	var applications []models.Application
	database.DB.Preload("Users").Find(&applications)
	json.NewEncoder(w).Encode(applications)
}

// Obtener todas las aplicaciones de un usuario
func GetApplicationsByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]

	var userApplications []models.UserApplication

	// Buscar todas las aplicaciones del usuario
	if err := database.DB.Where("user_id = ?", userID).Preload("Application").Find(&userApplications).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(userApplications)
}