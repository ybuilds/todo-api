package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func GetFromEnv(key string) string {
	return os.Getenv(key)
}
