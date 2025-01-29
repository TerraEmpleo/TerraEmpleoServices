package main

import (
    "log"
    "net/http"
    "os"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userProfileService/database"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userProfileService/routes"
    "github.com/gorilla/mux"
)

func main() {
	// Inicializar la base de datos
	database.InitDB()

	// Configurar rutas
	router := mux.NewRouter()
	routes.ProfileRoutes(router)

	// Obtener el puerto del .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// Iniciar el servidor
	log.Printf("Profile Service running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}