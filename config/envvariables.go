package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {

	envPath := "../.env"

	err := godotenv.Load(envPath)

	if err != nil {
		fmt.Print(err)
		log.Fatal("Error loading .env file!")
	}
}
