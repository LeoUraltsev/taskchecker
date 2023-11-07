package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	Http struct {
		Port string `env:"PORT" env-default:":8080"`
	}
	Log struct {
		Level string `env:"LOG_LEVEL" env-default:"debug"`
	}
}

var (
	config Config
	once   sync.Once
)

func New() *Config {

	once.Do(func() {
		if err := godotenv.Load(".env"); err != nil {
			log.Errorf("no .env file found: %v", err)
		}
		if err := cleanenv.ReadEnv(&config); err != nil {
			log.Fatalf("read config error: %v", err)
		}
	})

	return &config
}
