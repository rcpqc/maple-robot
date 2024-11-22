package scripts

import (
	"context"
	"time"
)

var container map[string]func(ctx context.Context)

func Register(name string, script func(ctx context.Context)) {
	container[name] = script
}

type Script struct {
	Name    string            `yaml:"name"`
	Record  time.Time         `yaml:"record"`
	Options map[string]string `yaml:"options"`
}
