package main

import (
    "log"
    "net/http"
    "os"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userService/database"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userService/routes"
    "github.com/gorilla/mux"
)

func main() {
    // Inicializar la base de datos
    database.InitDB()

    // Configurar rutas
    router := mux.NewRouter()
    routes.UserRoutes(router)

    // Iniciar el servidor
    port := os.Getenv("PORT")
    log.Printf("User Service running on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}
