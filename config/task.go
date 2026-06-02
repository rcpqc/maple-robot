package config

import (
	"context"
	"maple-robot/log"
	"sync"

	"github.com/rcpqc/expr"
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

func (o Condition) Match(vars expr.Vars) bool {
	if len(o) == 0 {
		return true
	}
	e, err := expr.Comp(string(o))
	if err != nil {
		panic(err)
	}
	return expr.EvalOr(e, vars, false).(bool)
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

// ---- disabled tasks (web toggle) ----

var (
	taskDisabledMu sync.RWMutex
	taskDisabled   = map[string]map[string]bool{} // scriptKey -> taskName -> disabled
)

func SetTaskDisabled(scriptKey, taskName string, disabled bool) {
	taskDisabledMu.Lock()
	if taskDisabled[scriptKey] == nil {
		taskDisabled[scriptKey] = map[string]bool{}
	}
	taskDisabled[scriptKey][taskName] = disabled
	taskDisabledMu.Unlock()
}

func IsTaskDisabled(scriptKey, taskName string) bool {
	taskDisabledMu.RLock()
	defer taskDisabledMu.RUnlock()
	if m := taskDisabled[scriptKey]; m != nil {
		return m[taskName]
	}
	return false
}

func AllDisabledTasks(scriptKey string) map[string]bool {
	taskDisabledMu.RLock()
	defer taskDisabledMu.RUnlock()
	out := map[string]bool{}
	for k, v := range taskDisabled[scriptKey] {
		out[k] = v
	}
	return out
}

// GetScriptTasks returns the task list for a script key.
func GetScriptTasks(key string) ([]*Task, error) {
	s, err := LoadScript(key)
	if err != nil {
		return nil, err
	}
	return s.Tasks, nil
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
