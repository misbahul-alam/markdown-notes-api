package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	AppPort   string `env:"APP_PORT,required"`
	JWTSecret string `env:"JWT_SECRET,required"`

	DBHost     string `env:"DB_HOST,required"`
	DBPort     string `env:"DB_PORT,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBName     string `env:"DB_NAME,required"`
	DBSSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`

	Env string `env:"ENV" envDefault:"development"`
}

var C Config

func Load() {
	if err := env.Parse(&C); err != nil {
		log.Fatalf("Config error %v", err)
	}

	log.Printf("Loaded config: %+v", C)
}
