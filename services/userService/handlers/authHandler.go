package handlers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userService/database"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userService/models"
)

// Validar si el rol proporcionado es válido
func isValidRole(role models.Role) bool {
	switch role {
	case models.RoleAdmin, models.RoleFarmer, models.RoleEmployer:
		return true
	default:
		return false
	}
}

// Registrar un usuario con validación de rol
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Si el usuario no proporciona un rol, asignar "employer" por defecto
	if user.Role == "" {
		user.Role = models.RoleEmployer
	}

	// Verificar si el rol es válido
	if !isValidRole(user.Role) {
		http.Error(w, "Invalid role. Must be 'admin', 'farmer', or 'employer'", http.StatusBadRequest)
		return
	}

	// Hashear la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error while hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Iniciar una transacción
	tx := database.DB.Begin()

	// Crear el usuario en la base de datos
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ✅ Crear un perfil vacío para el usuario recién creado
	profile := models.UserProfile{
		UserID:     user.ID,
		Location:   "",
		Skills:     "",
		Experience: 0,
		ResumeURL:  "",
		Bio:        "",
	}

	if err := tx.Create(&profile).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Error creating user profile", http.StatusInternalServerError)
		return
	}

	// Confirmar la transacción
	tx.Commit()

	// Enviar respuesta con el usuario y su perfil
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user":    user,
		"profile": profile,
	})
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
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Login successful",
		"user_id":  user.ID,
		"email":    user.Email,
	})
}

