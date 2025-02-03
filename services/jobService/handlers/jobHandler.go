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

