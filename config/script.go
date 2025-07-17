package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Script struct {
	Tasks []*Task `yaml:"tasks"`
}

func LoadScript(file string) (*Script, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	s := &Script{}
	if err := yaml.Unmarshal(bytes, s); err != nil {
		return nil, err
	}
	return s, nil
}
