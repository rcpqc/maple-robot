package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Web   WebConfig   `yaml:"web"`
	Adele *AdeleConfig `yaml:"adele"` // nil = 禁用隧道
	Auth  *AuthConfig `yaml:"auth"`   // nil = 不开启认证
}

type WebConfig struct {
	Addr string `yaml:"addr"`
}

type AuthConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type AdeleConfig struct {
	Server   string `yaml:"server"`   // 服务器地址 ip:port
	ClientID string `yaml:"client_id"` // 客户端标识
}

type Role struct {
	// Id     string `yaml:"id"`     //
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
