package config

import (
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Addr string `yaml:"addr"`
		Port string `yaml:"port"`
		Key  string `yaml:"key"`
		Cert string `yaml:"cert"`
	} `yaml:"server"`
	Rest struct {
		Addr  string `yaml:"addr"`
		Port  string `yaml:"port"`
		Root  string `yaml:"root"`
		Nodes struct {
			Hello string `yaml:"hello"`
		}
	}
}

var Conf *Config

func init() {
	l := echo.New().Logger
	c, err := Read()
	if err != nil {
		l.Fatalf("Config error: %s", err)
	}
	Conf = c
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

func Node(name string) string {
	s := BaseUrl()
	switch name {
	case "hello":
		return s + Conf.Rest.Nodes.Hello
	default:
		return s
	}
}

func BaseUrl() string {
	return strings.Join([]string{
		"http://",
		Conf.Rest.Addr,
		":",
		Conf.Rest.Port,
		Conf.Rest.Root,
	}, "")
}
