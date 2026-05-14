package config

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type ScriptsFile struct {
	Roles   []*Role           `yaml:"roles"`
	Scripts map[string]*Script `yaml:"scripts"`
}

type Script struct {
	Tasks []*Task `yaml:"tasks"`
}

var (
	sfOnce sync.Once
	sf     *ScriptsFile
	sfErr  error
)

func loadScriptsFile() (*ScriptsFile, error) {
	sfOnce.Do(func() {
		bytes, err := os.ReadFile("scripts.yaml")
		if err != nil {
			sfErr = fmt.Errorf("read scripts.yaml: %w", err)
			return
		}
		sf = &ScriptsFile{}
		if err := yaml.Unmarshal(bytes, sf); err != nil {
			sfErr = fmt.Errorf("parse scripts.yaml: %w", err)
			return
		}
	})
	return sf, sfErr
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
