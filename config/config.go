package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// Config holds application configuration
type Config struct {
	// Database settings
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBSSLMODE      string
	DBMaxOpenConns int // Maximum number of open connections to the database
	DBMaxIdleConns int // Maximum number of idle connections in the pool
	DBConnMaxLife  int // Maximum time (in seconds) a connection may be reused

	// API settings
	APIPort            string
	APIReadTimeout     int // Read timeout in seconds
	APIWriteTimeout    int // Write timeout in seconds
	APIIdleTimeout     int // Idle timeout in seconds
	APIShutdownTimeout int // Graceful shutdown timeout in seconds

	// Versioning
	APIVersion string // API version string
}

// LoadConfig reads environment variables and returns a Config struct
func LoadConfig() *Config {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg("No .env file found or error loading it. Using environment variables only.")
	}

	// Load configuration with defaults
	config := &Config{
		// Database defaults
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPassword:     getEnv("DB_PASSWORD", ""),
		DBName:         getEnv("DB_NAME", "goodschain"),
		DBSSLMODE:      getEnv("DB_SSL_MODE", "disable"),
		DBMaxOpenConns: getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
		DBMaxIdleConns: getEnvAsInt("DB_MAX_IDLE_CONNS", 5),
		DBConnMaxLife:  getEnvAsInt("DB_CONN_MAX_LIFE", 300), // 5 minutes

		// API defaults
		APIPort:            getEnv("API_PORT", "3000"),
		APIReadTimeout:     getEnvAsInt("API_READ_TIMEOUT", 15),
		APIWriteTimeout:    getEnvAsInt("API_WRITE_TIMEOUT", 15),
		APIIdleTimeout:     getEnvAsInt("API_IDLE_TIMEOUT", 60),
		APIShutdownTimeout: getEnvAsInt("API_SHUTDOWN_TIMEOUT", 30),

		// Versioning
		APIVersion: getEnv("API_VERSION", "v1"),
	}

	// Validate required configuration
	config.Validate()

	return config
}

// Validate validates the configuration and logs warnings/errors
func (c *Config) Validate() {
	// Critical configurations that must be present
	if c.DBName == "" {
		log.Fatal().Msg("Required configuration DB_NAME is missing")
	}

	// Check for empty credentials when not using localhost
	if c.DBHost != "localhost" && c.DBPassword == "" {
		log.Warn().Msg("Database password is empty for non-localhost database")
	}

	// API port validation
	if _, err := strconv.Atoi(c.APIPort); err != nil {
		log.Fatal().Err(err).Str("port", c.APIPort).Msg("Invalid API_PORT, must be a number")
	}

	// Log configuration (excluding sensitive data)
	log.Info().
		Str("db_host", c.DBHost).
		Str("db_port", c.DBPort).
		Str("db_name", c.DBName).
		Str("db_ssl_mode", c.DBSSLMODE).
		Int("db_max_open_conns", c.DBMaxOpenConns).
		Int("db_max_idle_conns", c.DBMaxIdleConns).
		Str("api_port", c.APIPort).
		Int("api_read_timeout", c.APIReadTimeout).
		Int("api_write_timeout", c.APIWriteTimeout).
		Str("api_version", c.APIVersion).
		Msg("Configuration loaded")
}

// GetDSN returns a formatted connection string for the database
func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMODE)
}

// Helper function to get an environment variable or return a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Helper function to get an environment variable as an integer or return a default value
func getEnvAsInt(key string, defaultValue int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
		log.Warn().Str("key", key).Msg("Invalid integer value in environment variable, using default")
	}
	return defaultValue
}
