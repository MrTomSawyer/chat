package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server Server
		DB     DB
	}

	Server struct {
		Addr string
	}

	DB struct {
		DSN string
	}
)

func MustLoadConfig(path string) *Config {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("failed to read config file: %s", err.Error()))
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Sprintf("failed to parse config file: %s", err.Error()))
	}

	return &config
}
