package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Env      string   `yaml:"env"`
	ExDelay  int64    `yaml:"external_delay"`
	InDelay  int64    `yaml:"internal_delay"`
	Telegram Telegram `yaml:"telegram"`
	HTTP     HTTP     `yaml:"http"`
	Postgres Postgres `yaml:"postgres"`
}

type Telegram struct {
	Token string `yaml:"bot_token"`
}

type HTTP struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"db_name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func MustLoad() *Config {
	var config Config

	path := os.Getenv("NOTIFICATION_SERVICE_CONFIG_PATH")
	if path == "" {
		log.Fatal("NOTIFICATION_SERVICE_CONFIG_PATH environment variable not set")
	}

	if _, err := os.ReadFile(path); err != nil {
		log.Fatalln(err)
	}

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		log.Fatalln(err)
	}

	return &config
}
