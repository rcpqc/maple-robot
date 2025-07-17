package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Roles []*Role `yaml:"roles"`
}

func Load(file string) (*Config, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
