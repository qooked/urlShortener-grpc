package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Env        string     `yaml:"env" default:"prod"`
	Url        string     `yaml:"url" env-required:"true"`
	Postgres   Postgres   `yaml:"postgres" env-required:"true"`
	GRPCserver GRPCserver `yaml:"gRPC" env-required:"true"`
	Redis      Redis      `yaml:"redis" env-required:"true"`
}

type GRPCserver struct {
	Timeout string `yaml:"timeout" env-default:"5s"`
	Host    string `yaml:"host" env-required:"true"`
	Port    int    `yaml:"port" env-default:"50051"`
}

type Postgres struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     int    `yaml:"port" env-required:"true"`
	Database string `yaml:"database" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
}

type Redis struct {
	Host string `yaml:"host" env-required:"true"`
	Port int    `yaml:"port" env-required:"true"`
}

func Load() (*Config, error) {
	var cfg *Config

	cfgByteArray, err := os.ReadFile("./config/cfg.yaml")
	if err != nil {
		err = fmt.Errorf("os.ReadFile(...): %w", err)
		return nil, err
	}

	err = yaml.Unmarshal(cfgByteArray, &cfg)
	if err != nil {
		err = fmt.Errorf("yaml.Unmarshal(...): %w", err)
		return nil, err
	}
	return cfg, nil
}
