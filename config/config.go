package config

import (
	"os"
)

type Config struct {
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	ServerPort     string
	AgifyURL       string
	GenderizeURL   string
	NationalizeURL string
}

func Load() *Config {
	return &Config{
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		ServerPort:     os.Getenv("SERVER_PORT"),
		AgifyURL:       os.Getenv("AGE_API_DOMAIN"),
		GenderizeURL:   os.Getenv("GENDER_API_DOMAIN"),
		NationalizeURL: os.Getenv("NATIONALITY_API_DOMAIN"),
	}
}
