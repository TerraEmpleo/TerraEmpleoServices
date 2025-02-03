package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userProfileService/database"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userProfileService/models"
	"strconv"
	fmt "fmt"
)

// Obtener todos los perfiles
func GetProfiles(w http.ResponseWriter, r *http.Request) {
	var profiles []models.UserProfile
	database.DB.Find(&profiles)
	json.NewEncoder(w).Encode(profiles)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userIDStr, exists := params["user_id"]
	if !exists || userIDStr == "" {
		http.Error(w, "User ID is required in the URL", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	var profile models.UserProfile
	if err := database.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
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

	// 🔹 Imprimir los parámetros recibidos en la URL para depuración
	fmt.Println("Parámetros recibidos:", params)

	userIDStr, exists := params["user_id"]
	if !exists || userIDStr == "" {
		http.Error(w, "User ID is required in the URL", http.StatusBadRequest)
		return
	}

	// 🔹 Convertir `user_id` a número
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	fmt.Println("UserID recibido:", userID) // ✅ Depuración

	var profile models.UserProfile

	// 🔹 Buscar el perfil por `UserID`
	if err := database.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}

	// 🔹 Decodificar el JSON del request
	var updatedProfile models.UserProfile
	if err := json.NewDecoder(r.Body).Decode(&updatedProfile); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// 🔹 Actualizar los campos
	profile.Location = updatedProfile.Location
	profile.Skills = updatedProfile.Skills
	profile.Experience = updatedProfile.Experience
	profile.ResumeURL = updatedProfile.ResumeURL
	profile.Bio = updatedProfile.Bio

	// 🔹 Guardar los cambios
	if err := database.DB.Save(&profile).Error; err != nil {
		http.Error(w, "Error updating profile", http.StatusInternalServerError)
		return
	}

	// 🔹 Responder con el perfil actualizado
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
