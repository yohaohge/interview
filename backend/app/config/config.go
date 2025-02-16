package config

import (
	"gopkg.in/ini.v1"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() (*Config, error) {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		return nil, err
	}

	return &Config{
		DBHost:     cfg.Section("database").Key("host").String(),
		DBPort:     cfg.Section("database").Key("port").String(),
		DBUser:     cfg.Section("database").Key("user").String(),
		DBPassword: cfg.Section("database").Key("password").String(),
		DBName:     cfg.Section("database").Key("name").String(),
	}, nil
}
