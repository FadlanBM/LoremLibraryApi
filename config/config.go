package config

import (
	"github.com/spf13/viper"
	"log"
)

type ConfigENV struct {
	PORT        string
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_DATABASE string
	DB_PORT     string
}

var ENV ConfigENV

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error loading config")
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatal(err)
	}

	log.Println("Config Env loaded")
}
