package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	// Server configuration
	ServerPort string
	ServerHost string

	// Database configuration
	Database DatabaseConfig

	// JWT configuration (for future implementation)
	JWT JWTConfig

	// Environment
	Environment string // "development" or "production"
}

// DatabaseConfig holds database connection settings
type DatabaseConfig struct {
	// Type: "sqlite" or "postgres"
	Type string

	// SQLite specific
	Path string // Path to SQLite database file

	// PostgreSQL specific
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// JWTConfig holds JWT authentication settings
type JWTConfig struct {
	Secret     string // Secret key for signing tokens
	Expiration int    // Token expiration time in hours
}

// Load reads configuration from environment variables
// It loads .env file if it exists, reads environment variables,
// and sets default values for development
func Load() *Config {
	// Load .env file (ignore error if file doesn't exist)
	_ = godotenv.Load()

	// Get environment (default: development)
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// Get server port (default: 8080)
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	// Get server host (default: localhost)
	serverHost := os.Getenv("SERVER_HOST")
	if serverHost == "" {
		serverHost = "localhost"
	}

	// Get database type (default: sqlite)
	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		dbType = "sqlite"
	}

	// Get database path (default: ./lottery.db)
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./lottery.db"
	}

	// Get JWT secret (default: empty, should be set in production)
	jwtSecret := os.Getenv("JWT_SECRET")

	// Get JWT expiration (default: 24 hours)
	jwtExpiration := 24
	if expStr := os.Getenv("JWT_EXPIRATION_HOURS"); expStr != "" {
		if exp, err := strconv.Atoi(expStr); err == nil {
			jwtExpiration = exp
		}
	}

	// PostgreSQL configuration (only used if DB_TYPE=postgres)
	dbConfig := DatabaseConfig{
		Type:     dbType,
		Path:     dbPath,
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  getEnvOrDefault("DB_SSLMODE", "disable"),
	}

	return &Config{
		ServerPort:  serverPort,
		ServerHost:  serverHost,
		Database:    dbConfig,
		Environment: env,
		JWT: JWTConfig{
			Secret:     jwtSecret,
			Expiration: jwtExpiration,
		},
	}
}

// GetDatabaseDSN returns the database connection string
// based on the database type (SQLite or PostgreSQL)
func (c *Config) GetDatabaseDSN() string {
	if c.Database.Type == "postgres" {
		// PostgreSQL connection string
		return fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			c.Database.Host,
			c.Database.Port,
			c.Database.User,
			c.Database.Password,
			c.Database.DBName,
			c.Database.SSLMode,
		)
	}
	// SQLite connection string (just the file path)
	return c.Database.Path
}

// getEnvOrDefault returns the environment variable value or a default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
