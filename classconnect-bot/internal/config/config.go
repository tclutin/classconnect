package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Token string `yaml:"token"`
	API   API    `yaml:"api"`
}

type API struct {
	BaseURL  string `yaml:"base_url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func MustLoad() *Config {
	var config Config

	path := os.Getenv("BOT_API_CONFIG_PATH")
	if path == "" {
		log.Fatalln("specify the correct path to the config")
	}

	if _, err := os.ReadFile(path); err != nil {
		log.Fatalln(err)
	}

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		log.Fatalln(err)
	}

	return &config
}
