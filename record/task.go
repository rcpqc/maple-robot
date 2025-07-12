package record

import (
	"context"
	"maple-robot/log"
	"strconv"
	"strings"
	"time"
)

var tasks = map[string]func(ctx context.Context, c *Context){}

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

func ProvideTask(name string, handler func(ctx context.Context, c *Context)) {
	tasks[name] = handler
}

func ExecuteTask(ctx context.Context, c *Context, name string) {
	if handler := tasks[name]; handler != nil {
		log.Info(ctx, "任务开始", "task", name)
		handler(ctx, c)
		log.Info(ctx, "任务完成", "task", name)
		c.Complete()
	} else {
		log.Error(ctx, "任务缺失", "task", name)
	}
}
