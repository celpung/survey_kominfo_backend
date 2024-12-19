package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	BASE_URL string
	PORT     string
}

var Env Environment

func init() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using hardcoded values")
	}

	// Initialize environment variables with fallback to hardcoded defaults
	Env = Environment{
		BASE_URL: getEnv("BASE_URL", "http://localhost:8080"),
		PORT:     getEnv("PORT", "8080"),
	}
}

// getEnv retrieves an environment variable or returns a default value if not set
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
