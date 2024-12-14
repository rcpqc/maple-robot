package context

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Load(file string) (*Context, error) {
	bytes, _ := os.ReadFile(file)
	ctx := &Context{file: file}
	if err := yaml.Unmarshal(bytes, ctx); err != nil {
		return nil, err
	}
	return ctx, nil
}

func (o *Context) Save() {
	bytes, _ := yaml.Marshal(o)
	os.WriteFile(o.file, bytes, 0644)
}

type Context struct {
	Roles []*Role `yaml:"roles"`

	file    string `yaml:"-"`
	curRole *Role  `yaml:"-"`
	curTask *Task  `yaml:"-"`
}

func (o *Context) Schedule() {
	o.curRole.Records.Mark(o.curTask.Name)
	o.Save()
}

func (o *Context) Complete() {
	o.curRole.Records.Mark(o.curTask.Name)
	o.Save()
}

func (o *Context) Lanuch() {
	o.curRole.Records.Start("角色")
}

func (o *Context) Finish() {
	o.curRole.Records.Mark("角色")
	o.Save()
}

func (o *Context) Execute(tasks []*Task) {
	for _, task := range tasks {
		if !task.Condition.Match() {
			continue
		}
		o.curTask = task
		// 脚本已调度，则跳过
		if o.curRole.Records.DailyDone(task.Name) {
			continue
		}
		o.curRole.Records.Start(task.Name)
		ExecuteTask(o, task.Name)
	}
}

func (o *Context) ExecuteSubs() {
	backupTask := o.curTask
	o.Execute(o.curTask.Subs)
	o.curTask = backupTask
}

func (o *Context) GetOption(name string) string {
	return o.curTask.Options[name]
}

func (o *Context) GetRoleIndex() int {
	return o.curRole.Index
}

func (o *Context) SetRole(role *Role) {
	o.curRole = role
}

func (o *Context) SetTask(task *Task) {
	o.curTask = task
}
