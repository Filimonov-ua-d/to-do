package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	viper.AutomaticEnv()
	return nil
}
