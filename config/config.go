package config

import (
	"errors"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	DB     DB     `yaml:"database"`
	Server Server `yaml:"server"`
}

type DB struct {
	PingDbAttempts int    `yaml:"pingdbattempts" env:"PING_DB_ATTEMPTS" env-default:"3" env-description:"tries to attempt to ping db"`
	SQLDSN         string `yaml:"sqldsn" env:"SQLDSN" env-default:"root:secret@tcp(morethanjustlinks-maria-db)/morethanjustlinks_db?charset=utf8mb4&parseTime=True&loc=Local" env-description:"sql dsn"`
}

type Server struct {
	Sessions []byte `yaml:"session" env:"SESSION" env-default:"secret" env-description:"session secret"`
}

func LoadConfig(configPath string) (AppConfig, error) {
	var config AppConfig

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		readEnvErr := cleanenv.ReadEnv(&config)
		if readEnvErr != nil {
			return config, errors.New("unable to read default environment variables")

		}
		return config, nil
	}

	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		return config, err
	}
	return config, nil

	// TODO check for required variables
}
