package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   Dbconfig
	Auth AuthConfig
}

type Dbconfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}

	dsn := os.Getenv("DSN")
	token := os.Getenv("TOKEN")

	log.Printf("Loaded DSN: %s", dsn)
	log.Printf("Loaded TOKEN: %s", token)

	return &Config{
		Db: Dbconfig{
			Dsn: dsn,
		},
		Auth: AuthConfig{
			Secret: token,
		},
	}
}
