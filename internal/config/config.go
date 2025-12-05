package config
import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


type Config struct {
	Port      string
	DbURL     string
	JWTSecret string
	Env       string
}

var AppConfig Config

func Load() {
	// Load .env file (optional in production)
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found, using system env variables")
	}

	AppConfig = Config{
		Port:      getEnv("PORT", "8080"),
		DbURL:     getEnv("DB_URL", ""),
		JWTSecret: getEnv("JWT_SECRET", ""),
		Env:       getEnv("ENV", "development"),
	}
}

// helper function with default value
func getEnv(key, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal
	}
	return value
}	