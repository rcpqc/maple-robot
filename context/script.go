package context

import (
	"maple-robot/log"
	"time"
)

var scripts = map[string]func(ctx *Context){}

type ScriptConfig struct {
	Name    string            `yaml:"name"`
	Options map[string]string `yaml:"options,omitempty"`
	Subs    []*ScriptConfig   `yaml:"subs,omitempty"`
	// 记录
	ScheduleTime time.Time `yaml:"schedule_time"`
	CompleteTime time.Time `yaml:"complete_time"`
}

func ProvideScript(name string, handler func(ctx *Context)) {
	scripts[name] = handler
}

func ExecuteScript(ctx *Context, name string) {
	if handler := scripts[name]; handler != nil {
		handler(ctx)
		ctx.Complete()
	} else {
		log.Warnf("script(%s) not provided", name)
	}
}
