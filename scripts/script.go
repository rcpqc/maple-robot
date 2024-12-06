package scripts

import (
	"maple-robot/context"
	"os"

	"gopkg.in/yaml.v3"
)

type Script struct {
	Tasks []*context.Task `yaml:"tasks"`
}

func Load(file string) (*Script, error) {
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
