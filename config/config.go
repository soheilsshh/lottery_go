package config

// TODO: Add imports when implementing
// "os"
// "github.com/joho/godotenv"

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
// It should:
// 1. Load .env file if it exists (using godotenv)
// 2. Read environment variables
// 3. Set default values for development
// 4. Return Config struct with all settings
func Load() *Config {
	// TODO: Implement configuration loading
	// - Load .env file
	// - Read DB_TYPE, DB_PATH, SERVER_PORT, JWT_SECRET, etc.
	// - Return populated Config struct
	return nil
}

// GetDatabaseDSN returns the database connection string
// based on the database type (SQLite or PostgreSQL)
func (c *Config) GetDatabaseDSN() string {
	// TODO: Implement DSN generation
	// - For SQLite: return file path
	// - For PostgreSQL: return connection string like "host=... port=... user=... password=... dbname=... sslmode=..."
	return ""
}
