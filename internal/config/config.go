package config

import (
	"flag"
	"os"
	"time"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	Debug    bool           `yaml:"debug"`
	Problem  ProblemConfig  `yaml:"problem"`
	Submit   SubmitConfig   `yaml:"submit"`
	Database DatabaseConfig `yaml:"database"`
	Server   ServerConfig   `yaml:"server"`
	Static   StaticConfig   `yaml:"static"`
}

type ProblemConfig struct {
	BasePath        string `yaml:"path"`
	ExplanationName string `yaml:"explanation"`
	InCasePrefix    string `yaml:"case_prefix_in"`
	InCaseSuffix    string `yaml:"case_suffix_in"`
	OutCasePrefix   string `yaml:"case_prefix_out"`
	OutCaseSuffix   string `yaml:"case_suffix_out"`
}

type SubmitConfig struct {
	BasePath             string `yaml:"path"`
	SocketPath           string `yaml:"socket"`
	MaxQueueSize         int32  `yaml:"maxQueueSize"`
	MaxConcurrentJudge   int32  `yaml:"maxConcurrent"`
	PJCompilerTypePrefix string `yaml:"compiler_type_prefix"`
	PJDetailName         string `yaml:"detail_name"`
}

type DatabaseConfig struct {
	DSN             string        `yaml:"dsn"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
}

type StaticConfig struct {
	BasePath string `yaml:"path"`
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

func LoadConfig(configPath string) (*Config, error) {
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
