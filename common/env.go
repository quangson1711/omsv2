package common

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key, fallback string) string {
	// Tải các biến môi trường từ tệp .env
	err := godotenv.Load("gateway/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
