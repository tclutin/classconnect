package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
)

const (
	prod string = "prod"
	dev  string = "dev"
)

type Config struct {
	Environment string `env:"ENV"`
	HTTPServer  HTTPServer
	Postgres    Postgres
	JWT         JWT
}

type HTTPServer struct {
	Address string `env:"HTTP_HOST" env-default:"localhost"`
	Port    string `env:"HTTP_PORT" env-default:"8080"`
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	DbName   string `env:"POSTGRES_DB"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
}

type JWT struct {
	Secret string `env:"JWT_SECRET"`
	Expire string `env:"JWT_EXPIRE"`
}

func MustLoad() *Config {
	var config Config

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	if err := cleanenv.ReadEnv(&config); err != nil {
		log.Fatal(err)
	}

	return &config
}

func (c *Config) IsProd() bool {
	return c.Environment == prod
}

func (c *Config) IsDev() bool {
	return c.Environment == dev
}
