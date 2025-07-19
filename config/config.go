package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Roles []*Role `yaml:"roles"`
}

type Role struct {
	Id     string `yaml:"id"`     //
	Class  string `yaml:"class"`  // 职业
	Script string `yaml:"script"` // 脚本
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
