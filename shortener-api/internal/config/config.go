package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Env         string      `yaml:"env" default:"prod"`
	DBstring    string      `yaml:"DBstring" env-required:"true"`
	Url         string      `yaml:"url" env-required:"true"`
	HttpPort    int         `yaml:"http-port" env-required:"true"`
	GRPCConfig  GRPCConfig  `yaml:"gRPC" env-required:"true"`
	RedisConfig RedisConfig `yaml:"redis" env-required:"true"`
}

type GRPCConfig struct {
	Timeout string `yaml:"timeout" env-default:"5s"`
	Port    int    `yaml:"port" env-default:"50051"`
}

type RedisConfig struct {
	Host string `yaml:"host" env-required:"true"`
	Port int    `yaml:"port" env-required:"true"`
}

var CFG *Config

func Get() {
	cfgByteArray, err := os.ReadFile("internal/config/cfg.yaml")

	if err != nil {
		panic("Error reading config file: " + err.Error())
	}

	err = yaml.Unmarshal(cfgByteArray, &CFG)

	if err != nil {
		panic("Error unmarshalling config file: " + err.Error())
	}
}
