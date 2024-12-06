package context

import (
	"maple-robot/log"
)

var tasks = map[string]func(ctx *Context){}

type Task struct {
	Name    string            `yaml:"name"`
	Options map[string]string `yaml:"options,omitempty"`
	Subs    []*Task           `yaml:"subs,omitempty"`
}

func ProvideTask(name string, handler func(ctx *Context)) {
	tasks[name] = handler
}

func ExecuteTask(ctx *Context, name string) {
	if handler := tasks[name]; handler != nil {
		handler(ctx)
		ctx.Complete()
	} else {
		log.Warnf("task(%s) not provided", name)
	}
}
