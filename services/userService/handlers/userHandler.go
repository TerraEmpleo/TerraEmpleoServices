package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/TerraEmpleo/TerraEmpleoServices/services/userService/database"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userService/models"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Preload("Profile").Find(&users)
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Model(&user).Updates(user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Obtener user_id de la URL
	params := mux.Vars(r)
	userIDStr, exists := params["user_id"]
	if !exists || userIDStr == "" {
		http.Error(w, "User ID is required in the URL", http.StatusBadRequest)
		return
	}

	// Convertir user_id a uint
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	// Buscar el usuario por ID
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Eliminar usuario
	if err := database.DB.Delete(&user).Error; err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("User with ID %d deleted successfully", userID),
	})
}