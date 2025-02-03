package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userService/database"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userService/routes"
)

func main() {
	// Inicializar la base de datos
	database.InitDB()

	// Crear router
	router := mux.NewRouter()

	// Registrar rutas
	routes.UserRoutes(router)

	// Configurar el puerto (por defecto 8080 si no está en el entorno)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// ✅ Aplicar middleware CORS con Gorilla Handlers
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // ✅ Permitir todas las peticiones (ajústalo en producción)
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
