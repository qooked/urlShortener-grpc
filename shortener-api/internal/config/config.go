package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Env        string `yaml:"env" default:"prod"`
	DBstring   string `yaml:"DBstring" env-required:"true"`
	Url        string `yaml:"url" env-required:"true"`
	GRPCConfig `yaml:"gRPC" env-required:"true"`
}

type GRPCConfig struct {
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
