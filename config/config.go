package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	DB struct {
		PING_DB_ATTEMPTS int    `yaml:"pingDbTimeouts" env:"PING_DB_ATTEMPTS" env-default:"3" env-description:"tries to attempt to ping db"`
		SQLDSN           string `yaml:"sqldsn" env:"SQLDSN" env-default:"root:secret@tcp(morethanjustlinks-maria-db)/morethanjustlinks_db?charset=utf8mb4&parseTime=True&loc=Local" env-description:"sql dsn"`
	} `yaml:"db"`
	Server struct {
		Sessions []byte `yaml:"session" env:"SESSION" env-default:"secret" env-description:"session secret"`
	} `yaml:"server"`
}

func LoadConfig(configPath string) (AppConfig, error) {
	var config AppConfig

	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
