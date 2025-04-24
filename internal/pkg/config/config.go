package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var configDirection string = "config/config.yml"

type Config struct {
	ServerPort string  `yaml:"port"`
	JWTSecret string 	`yaml:"secret"`
	Database   Postgre `yaml:"database"`
}

type Postgre struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile(configDirection)

	if err != nil {
		return nil, err
	}
	var OutPut Config

	if err := yaml.Unmarshal(data, &OutPut); err != nil {
		return nil, err
	}

	return &OutPut, nil

}
