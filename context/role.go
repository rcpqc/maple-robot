package context

import "time"

type RoleConfig struct {
	Index   int             `yaml:"index"`
	Scripts []*ScriptConfig `yaml:"scripts"`
	// 记录
	LaunchTime time.Time `yaml:"launch_time"`
	FinishTime time.Time `yaml:"finish_time"`
}
