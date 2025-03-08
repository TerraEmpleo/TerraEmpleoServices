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
	var requestData struct {
		UserID uint   `json:"user_id"`
		JobID  uint   `json:"job_id"`
		Status string `json:"status"`
	}

	// Decodificar JSON del request
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Error en la solicitud: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar si el usuario existe
	var user models.User
	if err := database.DB.Where("id = ?", requestData.UserID).First(&user).Error; err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	// Verificar si el trabajo existe
	var job models.Job
	if err := database.DB.Where("id = ?", requestData.JobID).First(&job).Error; err != nil {
		http.Error(w, "Trabajo no encontrado", http.StatusNotFound)
		return
	}

	// Verificar si el usuario ya aplicó a este trabajo
	var existingApplication models.UserApplication
	if err := database.DB.Where("user_id = ? AND application_id IN (SELECT id FROM applications WHERE job_id = ?)", requestData.UserID, requestData.JobID).First(&existingApplication).Error; err == nil {
		http.Error(w, "El usuario ya ha aplicado a este trabajo", http.StatusConflict)
		return
	}

	// Crear la aplicación en la tabla `applications`
	newApplication := models.Application{
		JobID: requestData.JobID,
	}
	if err := database.DB.Create(&newApplication).Error; err != nil {
		http.Error(w, "Error al guardar la aplicación en la tabla applications: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Crear la relación en la tabla `user_applications`
	newUserApplication := models.UserApplication{
		UserID:        requestData.UserID,
		ApplicationID: newApplication.ID,
		Status:        requestData.Status,
	}
	if err := database.DB.Create(&newUserApplication).Error; err != nil {
		http.Error(w, "Error al guardar la relación en user_applications: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder con la información creada
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"application":      newApplication,
		"user_application": newUserApplication,
	})
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

    var applications []models.Application

    // Primero, obtener las aplicaciones en las que el usuario ha aplicado
    if err := database.DB.Where("id IN (SELECT application_id FROM user_applications WHERE user_id = ?)", userID).
        Find(&applications).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Luego, obtener los Jobs relacionados con estas aplicaciones
    var jobIDs []uint
    for _, application := range applications {
        jobIDs = append(jobIDs, application.JobID)
    }

    var jobs []models.Job
    if err := database.DB.Where("id IN ?", jobIDs).Find(&jobs).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respuesta con aplicaciones y sus trabajos relacionados
    response := map[string]interface{}{
        "applications": applications,
        "jobs":         jobs,
    }

    json.NewEncoder(w).Encode(response)
}


