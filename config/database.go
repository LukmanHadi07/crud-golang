package config

import (
	"fmt"
	"log"
	"mahasiswa-api-golang/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	// Ambil variabel dari environment Railway
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("PGDATABASE")

	// Pastikan semua variabel tersedia
	if host == "" || port == "" || user == "" || password == "" || dbName == "" {
		log.Fatal("One or more required database environment variables are missing")
	}

	// Buat DSN (Database Source Name)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	// Koneksi ke PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// AutoMigrate model Students
	err = db.AutoMigrate(&model.Students{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return db
}
