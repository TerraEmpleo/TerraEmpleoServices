package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/applicationService/database"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/applicationService/routes"
)

func main() {
	// Inicializar la base de datos
	database.InitDB()

	// Configurar rutas
	router := mux.NewRouter()
	routes.ApplicationRoutes(router)

	// Obtener el puerto del .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
	}

	// Iniciar el servidor
	log.Printf("Application Service running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
