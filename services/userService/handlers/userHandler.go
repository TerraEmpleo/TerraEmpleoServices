package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/TerraEmpleo/TerraEmpleoServices/services/userService/database"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userService/models"
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
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User deleted successfully",
	})
}
