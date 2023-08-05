package config

import (
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	HTTP struct {
		ListenAddress string `toml:"listen_address"`
		Debug         bool   `toml:"debug"`
	}
	Logger struct {
		LoggingLevel uint `toml:"level"`
	}
}

func (cfg *AppConfig) Init() error {
	_, err := toml.DecodeFile("config.toml", cfg)
	if err != nil {
		return err
	}

	err = godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}

func New() *AppConfig {
	return &AppConfig{}
}
