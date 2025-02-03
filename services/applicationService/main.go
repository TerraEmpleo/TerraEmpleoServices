package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/applicationService/database"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/applicationService/routes"
	"github.com/gorilla/handlers"
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

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // Métodos permitidos
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), // Headers permitidos
	)

	// Iniciar el servidor con CORS habilitado
	log.Printf("User Service running on port %s", port)
	err := http.ListenAndServe(":"+port, corsHandler(router))
	if err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
	}
}
