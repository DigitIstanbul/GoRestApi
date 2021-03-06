package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {

	path, _ := os.Getwd()

	if os.Getenv("APPLICATION_ENV") == "production" {
		err := godotenv.Load(path + "/env/production.env")
		if err != nil {
			log.Fatal("Error loading production.env file")
		}
	} else {
		err := godotenv.Load(path + "/env/development.env")
		if err != nil {
			log.Fatal("Error loading development.env file, " + path)
		}
	}
}

func GetConfig(key string) string {
	return os.Getenv(key)
}
