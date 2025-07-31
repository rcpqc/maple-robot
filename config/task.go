package config

import (
	"context"
	"maple-robot/log"
	"strconv"
	"strings"
	"time"
)

type taskkey struct{}

var tasks = map[string]func(ctx context.Context){}

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

func ProvideTask(name string, handler func(ctx context.Context)) {
	tasks[name] = handler
}

func WithTaskOptions(ctx context.Context, options map[string]string) context.Context {
	return context.WithValue(ctx, taskkey{}, options)
}

func GetTaskOptions(ctx context.Context, key string) string {
	options, ok := ctx.Value(taskkey{}).(map[string]string)
	if !ok {
		return ""
	}
	return options[key]
}

func (o *Task) Execute(ctx context.Context) {
	ctx = log.WithLogger(ctx, log.GetLogger(ctx).With("task", o.Name))
	if handler := tasks[o.Name]; handler != nil {
		log.Info(ctx, "任务开始")
		handler(WithTaskOptions(ctx, o.Options))
		log.Info(ctx, "任务完成")
	} else {
		log.Error(ctx, "任务缺失")
	}
}
