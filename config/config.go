package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Addr string `yaml:"addr"`
		Port string `yaml:"port"`
		Key  string `yaml:"key"`
		Cert string `yaml:"cert"`
	} `yaml:"server"`
}

func Read() (*Config, error) {
	cfg := &Config{}
	f, err := os.Open("config.yaml")
	if err != nil {
		return cfg, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, err
}
