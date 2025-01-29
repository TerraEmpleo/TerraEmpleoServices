package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userProfileService/database"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userProfileService/models"
)

// Obtener todos los perfiles
func GetProfiles(w http.ResponseWriter, r *http.Request) {
	var profiles []models.UserProfile
	database.DB.Find(&profiles)
	json.NewEncoder(w).Encode(profiles)
}

// Obtener un perfil por ID
func GetProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var profile models.UserProfile
	if err := database.DB.First(&profile, params["id"]).Error; err != nil {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(profile)
}

// Crear un nuevo perfil
func CreateProfile(w http.ResponseWriter, r *http.Request) {
	var profile models.UserProfile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&profile).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}

// Actualizar un perfil
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var profile models.UserProfile

	if err := database.DB.First(&profile, params["id"]).Error; err != nil {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Save(&profile).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(profile)
}

// Eliminar un perfil
func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var profile models.UserProfile

	if err := database.DB.First(&profile, params["id"]).Error; err != nil {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}

	database.DB.Delete(&profile)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Profile deleted successfully",
	})
}
