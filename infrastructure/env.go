package infrastructure

import (
	"log"

	"github.com/joho/godotenv"
)

func InitLoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
