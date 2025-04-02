package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Port string

var DBUser string
var DBPassword string
var DBName string
var DBHost string

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}

	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	DBHost = os.Getenv("DB_HOST")

	if DBUser == "" || DBPassword == "" || DBName == "" || DBHost == "" {
		log.Fatal("DB_USER, DB_PASSWORD, DB_NAME must be set in .env file")
	}

}
