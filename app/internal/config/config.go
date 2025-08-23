package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	LogLevel string `yaml:"log_level" env-default:"trace"`
	HttpServer struct {
		Host string `yaml:"host" env-default:"postgres"`
		Port string `yaml:"port" env-default:"8000"`
	} `yaml:"http_server"`
	Postgres struct {
		DB_HOST string `yaml:"DB_HOST" env-default:"localhost"`
		DB_PORT string `yaml:"DB_PORT" env-default:"5432"`
		DB_USER string `yaml:"DB_USER" env-default:"postgres"`
		DB_PASS string `yaml:"DB_PASSWORD" env-default:"password"`
		DB_NAME string `yaml:"DB_NAME" env-default:"rest_api"`
	}
}

var instance *Config
var once sync.Once

func MastLoad() *Config {
	once.Do(func() {
		instance = &Config{}

		configPath := filepath.Join("configs", "config.yaml")
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			panic("The configuration file was not found on:" + configPath)
		}

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			panic(err)
		}
	})
	return instance
}