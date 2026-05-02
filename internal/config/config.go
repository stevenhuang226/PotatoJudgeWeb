package config

import (
	"flag"
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	Debug    bool           `yaml:"debug"`
	Problem  ProblemConfig  `yaml:"problem"`
	Submit   SubmitConfig   `yaml:"submit"`
	Database DatabaseConfig `yaml:"database"`
	Server   ServerConfig   `yaml:"server"`
}

type ProblemConfig struct {
	Path string `yaml:"path"`
}

type SubmitConfig struct {
	Path               string `yaml:"path"`
	MaxConcurrentJudge uint16 `yaml:"maxConcurrent"`
}

type DatabaseConfig struct {
	Path string `yaml:"path"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
}

const defaultConfigPath string = "./config.yaml"

func resolveConfigPath() string {
	var configPath string

	flag.StringVar(&configPath, "config", "", "config file path")
	flag.Parse()

	if configPath != "" {
		return configPath
	}

	return defaultConfigPath
}

func load_config(configPath string) (*Config, error) {
	raw, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(raw, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
