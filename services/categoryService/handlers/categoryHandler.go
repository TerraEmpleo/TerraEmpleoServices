package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/categoryService/database"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/categoryService/models"
)

// Obtener todas las categorías
func GetCategories(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	database.DB.Find(&categories)
	json.NewEncoder(w).Encode(categories)
}

// Obtener una categoría por ID
func GetCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var category models.Category
	if err := database.DB.First(&category, params["id"]).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(category)
}

// Crear una nueva categoría
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&category).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

// Actualizar una categoría
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var category models.Category

	if err := database.DB.First(&category, params["id"]).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Save(&category).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(category)
}

// Eliminar una categoría
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var category models.Category

	if err := database.DB.First(&category, params["id"]).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	database.DB.Delete(&category)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Category deleted successfully",
	})
}
