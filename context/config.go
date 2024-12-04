package context

import (
	"os"
	"time"

	"maple-robot/log"

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
	Roles []*RoleConfig `yaml:"roles"`

	file      string        `yaml:"-"`
	curRole   *RoleConfig   `yaml:"-"`
	curScript *ScriptConfig `yaml:"-"`
}

func (o *Context) Schedule() {
	o.curScript.ScheduleTime = time.Now()
	o.Save()
}

func (o *Context) Complete() {
	o.curScript.CompleteTime = time.Now()
	o.Save()
}

func (o *Context) Lanuch() {
	o.curRole.LaunchTime = time.Now()
	o.Save()
}

func (o *Context) Finish() {
	o.curRole.FinishTime = time.Now()
	o.Save()
}

func (o *Context) Startup() {
	for _, role := range o.Roles {
		o.curRole = role
		// 角色已完成, 则跳过
		if role.FinishTime.Format("2006-01-02") == time.Now().Format("2006-01-02") {
			log.Infof("[角色-%d] 已完成", role.Index)
			continue
		}
		o.Lanuch()
		for _, script := range role.Scripts {
			o.curScript = script
			// 脚本已调度，则跳过
			if script.ScheduleTime.Format("2006-01-02") == time.Now().Format("2006-01-02") {
				log.Infof("[角色-%d] [脚本-%s] 已调度", role.Index, script.Name)
				continue
			}
			ExecuteScript(o, script.Name)
		}
		o.Finish()
	}
}

func (o *Context) GetOption(name string) string {
	return o.curScript.Options[name]
}

func (o *Context) GetRoleIndex() int {
	return o.curRole.Index
}

func (o *Context) SetRole(role *RoleConfig) {
	o.curRole = role
}

func (o *Context) SetScript(script *ScriptConfig) {
	o.curScript = script
}
