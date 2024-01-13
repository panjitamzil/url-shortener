package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App   AppConfig   `yaml:"app"`
	Redis RedisConfig `yaml:"redis"`
}

type AppConfig struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type RedisConfig struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

var Cfg Config

func LoadConfig(filename string) (err error) {
	configByte, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	yaml.Unmarshal(configByte, &Cfg)

	return
}
