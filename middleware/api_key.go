package middleware

import (
	"log"
	"os"
	"crypto/sha256"
	"crypto/subtle"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v3/middleware/keyauth"
)


func ValidateAPIKey(c fiber.Ctx, key string) (bool, error) {
	log.Println("Validating API Key")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	log.Println(apiKey)
	if apiKey == "" {
		log.Fatal("API_KEY is not set in the environment")
	}
	hashedAPIKey := sha256.Sum256([]byte(apiKey))
	hashedKey := sha256.Sum256([]byte(key))

	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

