package database

import (
	"fmt"
	"log"
	"os"
	"github.com/TerraEmpleo/TerraEmpleoServices/services/userService/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDB() {
	// Cargar las variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Construir la cadena de conexión
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	// Conectar a la base de datos
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrar los modelos
	err = DB.AutoMigrate(&models.User{}, &models.UserProfile{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database connected and migrated successfully!")
}
