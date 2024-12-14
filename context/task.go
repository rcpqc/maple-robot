package context

import (
	"maple-robot/log"
	"strconv"
	"strings"
	"time"
)

var tasks = map[string]func(ctx *Context){}

type Task struct {
	Name      string            `yaml:"name"`
	Condition Condition         `yaml:"condition"`
	Options   map[string]string `yaml:"options,omitempty"`
	Subs      []*Task           `yaml:"subs,omitempty"`
}

type Condition string

func (o Condition) Match() bool {
	if len(o) == 0 {
		return true
	}
	return strings.Contains(string(o), strconv.FormatInt(int64(time.Now().Weekday()), 10))
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
