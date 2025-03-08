package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"time"
	"github.com/gorilla/mux"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/jobService/database"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/jobService/models"
	"strconv"
	"log"
	"strings"
)

// Obtener todas las ofertas de empleo
func GetJobs(w http.ResponseWriter, r *http.Request) {
	var jobs []models.Job
	database.DB.Find(&jobs)
	json.NewEncoder(w).Encode(jobs)
}

// Obtener una oferta de empleo por ID
func GetJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var job models.Job
	if err := database.DB.First(&job, params["id"]).Error; err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(job)
}

// Crear una nueva oferta de empleo con validación de empleador
// Crear una nueva oferta de empleo con validación de empleador y cambio de rol
func CreateJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Obtener `user_id` de la URL
	userIDStr, exists := params["user_id"]
	if !exists || userIDStr == "" {
		http.Error(w, "User ID is required in the URL", http.StatusBadRequest)
		return
	}

	// Convertir `user_id` a número
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	// Verificar si el usuario existe
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		http.Error(w, fmt.Sprintf("User with ID %d not found", userID), http.StatusNotFound)
		return
	}

	// Si el usuario tiene el rol "farmer", actualizarlo a "employer"
	if user.Role == "farmer" {
		if err := database.DB.Model(&user).Update("role", "employer").Error; err != nil {
			http.Error(w, "Error updating user role", http.StatusInternalServerError)
			return
		}
	}

	// Decodificar el JSON del request
	var requestBody struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Location    string  `json:"location"`
		Salary      float64 `json:"salary"`
		CategoryID  uint    `json:"category_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 🟢 Imprimir valores para depuración
	log.Println("CategoryID recibido:", requestBody.CategoryID)

	// Crear la oferta con los valores recibidos
	job := models.Job{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		Location:    requestBody.Location,
		Salary:      requestBody.Salary,
		EmployerID:  uint(userID),
		CategoryID:  requestBody.CategoryID,
	}

	// Iniciar la transacción
	tx := database.DB.Begin()

	// Verificar que la categoría existe
	var category models.Category
	if err := tx.First(&category, job.CategoryID).Error; err != nil {
		log.Println("Error: Categoría no encontrada - ID:", job.CategoryID)
		tx.Rollback()
		http.Error(w, fmt.Sprintf("Category with ID %d does not exist", job.CategoryID), http.StatusBadRequest)
		return
	}

	// Crear la oferta de empleo dentro de la transacción
	if err := tx.Create(&job).Error; err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Confirmar la transacción
	tx.Commit()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(job)
}




// Actualizar una oferta de empleo
func UpdateJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var job models.Job

	// Verificar que el trabajo existe
	if err := database.DB.First(&job, params["id"]).Error; err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	// Decodificar el request en un nuevo job temporal
	var updatedJob models.Job
	if err := json.NewDecoder(r.Body).Decode(&updatedJob); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar que el EmployerID no se modifique
	if updatedJob.EmployerID != job.EmployerID {
		http.Error(w, "You cannot change the EmployerID of a job", http.StatusForbidden)
		return
	}

	// Actualizar solo los campos permitidos
	job.Title = updatedJob.Title
	job.Description = updatedJob.Description
	job.Location = updatedJob.Location
	job.Salary = updatedJob.Salary
	job.Requirements = updatedJob.Requirements
	job.UpdatedAt = time.Now()

	if err := database.DB.Save(&job).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(job)
}


// Eliminar una oferta de empleo
func DeleteJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var job models.Job

	if err := database.DB.First(&job, params["id"]).Error; err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	database.DB.Delete(&job)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Job deleted successfully",
	})
}
func SearchJobs(w http.ResponseWriter, r *http.Request) {
	var jobs []models.Job
	query := database.DB

	// Obtener parámetros de consulta (query params)
	title := r.URL.Query().Get("title")
	location := r.URL.Query().Get("location")
	minSalary := r.URL.Query().Get("min_salary")
	maxSalary := r.URL.Query().Get("max_salary")
	categoryID := r.URL.Query().Get("category_id")

	// Aplicar filtros dinámicos
	if title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}
	if location != "" {
		query = query.Where("location ILIKE ?", "%"+location+"%")
	}
	if minSalary != "" {
		query = query.Where("salary >= ?", minSalary)
	}
	if maxSalary != "" {
		query = query.Where("salary <= ?", maxSalary)
	}
	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	// Ejecutar consulta
	if err := query.Find(&jobs).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(jobs)
}
// Obtener los trabajos creados por un usuario específico
func GetJobsByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIDStr, exists := params["user_id"]

	if !exists || userIDStr == "" {
		http.Error(w, "User ID is required in the URL", http.StatusBadRequest)
		return
	}

	// Convertir `user_id` a uint
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	// Buscar trabajos donde EmployerID coincida con el `user_id`
	var jobs []models.Job
	if err := database.DB.Where("employer_id = ?", userID).Find(&jobs).Error; err != nil {
		http.Error(w, "Error retrieving jobs", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(jobs)
}


// Obtener recomendaciones de trabajos para un farmer basado en ubicación
func RecommendJobsByLocationForFarmer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIDStr, exists := params["user_id"]

	if !exists || userIDStr == "" {
		http.Error(w, "User ID is required in the URL", http.StatusBadRequest)
		return
	}

	// Convertir `user_id` a uint
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	// Obtener el perfil del `farmer`
	var profile models.UserProfile
	if err := database.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		http.Error(w, "User profile not found", http.StatusNotFound)
		return
	}

	// Verificar si el `farmer` tiene una ubicación registrada
	if profile.Location == "" {
		http.Error(w, "Farmer location is not set", http.StatusBadRequest)
		return
	}

	// Buscar trabajos en la misma ubicación
	var recommendedJobs []models.Job
	if err := database.DB.Where("location = ?", profile.Location).Order("created_at DESC").Limit(5).Find(&recommendedJobs).Error; err != nil {
		http.Error(w, "Error retrieving jobs", http.StatusInternalServerError)
		return
	}

	// Si no hay trabajos en la ubicación, sugerir los más recientes de cualquier ubicación
	if len(recommendedJobs) == 0 {
		database.DB.Order("created_at DESC").Limit(5).Find(&recommendedJobs)
	}

	// Enviar las recomendaciones
	json.NewEncoder(w).Encode(recommendedJobs)
}


// Obtener recomendaciones de `farmers` para un `employer`
func RecommendFarmersForEmployer(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)
	jobIDStr, exists := params["job_id"]
	log.Println("Job ID recibido como string:", jobIDStr)
	if !exists || jobIDStr == "" {
		http.Error(w, "Job ID is required in the URL", http.StatusBadRequest)
		return
	}

	// Convertir `job_id` a uint
	jobID, err := strconv.ParseUint(jobIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid Job ID", http.StatusBadRequest)
		return
	}

	// Obtener los detalles del `job`
	var job models.Job
	if err := database.DB.First(&job, jobID).Error; err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	log.Println("Buscando farmers para el job:", job.Title, "en", job.Location)

	// Buscar `farmers` en la misma ubicación y con habilidades relevantes
	var farmers []models.UserProfile
	query := database.DB.Where("location = ?", job.Location)

	// Si el `job` tiene requisitos, intentamos buscar `farmers` con esas habilidades
	if job.Requirements != "" {
		skills := strings.Split(job.Requirements, ",") // Dividir las habilidades en una lista
		for _, skill := range skills {
			query = query.Or("skills ILIKE ?", "%"+strings.TrimSpace(skill)+"%")
		}
	}

	// Obtener `farmers` recomendados
	err = query.Order("experience DESC").Limit(5).Find(&farmers).Error
	if err != nil {
		http.Error(w, "Error retrieving farmers", http.StatusInternalServerError)
		return
	}

	// Si no hay `farmers` en la ubicación, recomendar los más experimentados sin importar la ubicación
	if len(farmers) == 0 {
		database.DB.Order("experience DESC").Limit(5).Find(&farmers)
	}

	// Enviar las recomendaciones
	json.NewEncoder(w).Encode(farmers)
}
