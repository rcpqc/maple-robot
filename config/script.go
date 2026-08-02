package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ScriptsFile struct {
	Roles   []*Role            `yaml:"roles"`
	Scripts map[string]*Script `yaml:"scripts"`
}

type Script struct {
	Tasks []*Task `yaml:"tasks"`
}

func loadScriptsFile() (*ScriptsFile, error) {
	bytes, err := os.ReadFile("scripts.yaml")
	if err != nil {
		return nil, fmt.Errorf("read scripts.yaml: %w", err)
	}
	sf := &ScriptsFile{}
	if err := yaml.Unmarshal(bytes, sf); err != nil {
		return nil, fmt.Errorf("parse scripts.yaml: %w", err)
	}
	return sf, nil
}

// LoadRoles 返回 scripts.yaml 中的角色列表.
func LoadRoles() ([]*Role, error) {
	f, err := loadScriptsFile()
	if err != nil {
		return nil, err
	}
	return f.Roles, nil
}

// LoadScript 从 scripts.yaml 中按 key 加载脚本.
func LoadScript(key string) (*Script, error) {
	if key == "" {
		key = "200"
	}
	f, err := loadScriptsFile()
	if err != nil {
		return nil, err
	}
	s, ok := f.Scripts[key]
	if !ok {
		return nil, fmt.Errorf("script key %q not found in scripts.yaml", key)
	}
	return s, nil
}
