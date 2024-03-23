package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Env  string `yaml:"env"`
	Port int    `yaml:"port" env-required:"true"`
	Addr string `yaml:"addr" env-required:"true"`
	GRPC GRPC   `yaml:"gRPC"`
}

type GRPC struct {
	Port    int    `yaml:"port" env-default:"50051"`
	Timeout string `yaml:"timeout" env-default:"5s"`
}

func Get() *Config {
	cfgFile, err := os.ReadFile("./internal/config/cfg.yaml")

	if err != nil {
		panic("Error reading config file: " + err.Error())
	}

	var cfg Config
	err = yaml.Unmarshal(cfgFile, &cfg)

	if err != nil {
		panic("Error unmarshalling config file: " + err.Error())
	}

	slog.Info("Loaded config file")
	return &cfg
}
