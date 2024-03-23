package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Env        string `yaml:"env" default:"prod"`
	DBstring   string `yaml:"DBstring" env-required:"true"`
	BotToken   string `yaml:"bot" env-required:"true"`
	GRPCConfig `yaml:"gRPC" env-required:"true"`
	Clients    ClientsConfig `yaml:"clients"`
}

type GRPCConfig struct {
	Timeout string `yaml:"timeout" env-default:"5s"`
	Port    int    `yaml:"port" env-default:"50051"`
}

type Client struct {
	Adress       string        `yaml:"adress"`
	Timeout      time.Duration `yaml:"timeout"`
	RetriesCount int           `yaml:"retries"`
}

type ClientsConfig struct {
	API Client `yaml:"api"`
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
