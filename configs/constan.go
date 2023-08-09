package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetPathAPI() string {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}
	result := os.Getenv("PATH_API")
	log.Println(result)
	return result
}
