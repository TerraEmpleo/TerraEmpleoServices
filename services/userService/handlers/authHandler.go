package handlers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userService/database"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userService/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hashear la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error while hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Crear el usuario en la base de datos
	if err := database.DB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Buscar al usuario por email
	var user models.User
	if err := database.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Verificar la contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
	})
}
