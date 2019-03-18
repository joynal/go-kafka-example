package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfigs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
