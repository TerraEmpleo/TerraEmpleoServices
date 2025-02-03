package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/TerraEmpleo/TerraEmpleoServices/db/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "github.com/lib/pq" // Para usar con sql.Open
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Obtener datos básicos para la conexión inicial
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	dbTimeZone := os.Getenv("DB_TIMEZONE")

	// Crear conexión básica al servidor (sin especificar una base de datos específica)
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPassword, dbPort, dbSSLMode, dbTimeZone,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL server: %v", err)
	}
	defer db.Close()

	// Verificar si la base de datos existe
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	if err != nil {
		log.Printf("Database '%s' already exists or cannot be created: %v", dbName, err)
	} else {
		log.Printf("Database '%s' created successfully.", dbName)
	}

	// Conectar a la base de datos específica
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode, dbTimeZone,
	)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrar las tablas
	err = gormDB.AutoMigrate(
		&models.User{},
		&models.UserProfile{},
		&models.Job{},
		&models.Application{},
		&models.Recommendation{},
		&models.Category{},
		&models.Feedback{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migrated successfully!")
}