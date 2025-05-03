package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config holds application configuration
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMODE  string
	APIPort    string
}

// LoadConfig reads environment variables and returns a Config struct
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMODE:  os.Getenv("DB_SSL_MODE"),
		APIPort:    os.Getenv("API_PORT"),
	}
}
