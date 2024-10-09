package db

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "anigo/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    dbHost := "localhost"
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")

    dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    DB.AutoMigrate(&models.Anime{})
}
