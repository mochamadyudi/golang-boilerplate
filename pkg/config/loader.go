package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error when loading file configuration", err.Error())
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Pass:     os.Getenv("DB_PASS"),
			TimeZone: os.Getenv("DB_TIMEZONE"),
			SSlMode:  os.Getenv("DB_SSL_MODE"),
		},
		Jwt: JsonWebToken{
			Secret:      os.Getenv("JWT_SECRET"),
			ExpiresHour: os.Getenv("JWT_EXPIRES_HOUR"),
		},
	}

}
