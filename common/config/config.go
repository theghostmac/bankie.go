package config

import "os"

type Config struct {
	Env  string `envconfig:"env"`
	Port int    `envconfig:"http_port"`
}

func InitiateConfig(config *Config, err error) {
	// Initiate environment variables.
	if _, err := os.Stat(".env"); err != nil {
		if !os.IsNotExist(err) {
			//return nil, err
		}
	} else {
	}
}
