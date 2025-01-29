package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/categoryService/database"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/categoryService/routes"
)

func main() {
	// Inicializar la base de datos
	database.InitDB()

	// Configurar rutas
	router := mux.NewRouter()
	routes.CategoryRoutes(router)

	// Obtener el puerto del .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	// Iniciar el servidor
	log.Printf("Category Service running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
