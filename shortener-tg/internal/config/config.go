package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Env      string `yaml:"env" default:"prod"`
	Addr     string `yaml:"addr" env-required:"true"`
	BotToken string `yaml:"bot" env-required:"true"`
	GRPC     GRPC   `yaml:"gRPC" env-required:"true"`
}

type GRPC struct {
	Timeout string `yaml:"timeout" env-default:"5s"`
	Port    int    `yaml:"port" env-default:"50051"`
}

func Get() *Config {
	cfgByteArray, err := os.ReadFile("internal/config/cfg.yaml")

	if err != nil {
		panic("Error reading config file: " + err.Error())
	}

	var cfg Config
	err = yaml.Unmarshal(cfgByteArray, &cfg)

	if err != nil {
		panic("Error unmarshalling config file: " + err.Error())
	}

	return &cfg
}
