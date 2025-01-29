package main

import (
    "log"
    "net/http"
    "os"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/jobService/database"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/jobService/routes"
    "github.com/gorilla/mux"
)

func main() {
	// Inicializar la base de datos
	database.InitDB()

	// Configurar rutas
	router := mux.NewRouter()
	routes.JobRoutes(router)

	// Obtener el puerto del .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	// Iniciar el servidor
	log.Printf("Jobs Service running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}