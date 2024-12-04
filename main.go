package main

import (
	"maple-robot/context"
	"maple-robot/log"
	"maple-robot/scripts"
	"time"
)

func gen() {
	ctx, _ := context.Load("context.yaml")
	ctx.Roles = nil
	for page := 1; page <= 4; page++ {
		for pos := 1; pos <= 7; pos++ {
			t := time.Now().AddDate(0, 0, -1)
			if page < 3 || (page == 3 && pos < 2) {
				t = time.Now()
			}
			role := &context.RoleConfig{Index: page*100 + pos, LaunchTime: t, FinishTime: t}
			role.Scripts = append(role.Scripts, &context.ScriptConfig{Name: "自动战斗时间", ScheduleTime: t, CompleteTime: t})
			role.Scripts = append(role.Scripts, &context.ScriptConfig{Name: "补充训练时间", ScheduleTime: t, CompleteTime: t})
			role.Scripts = append(role.Scripts, &context.ScriptConfig{Name: "送人气", ScheduleTime: t, CompleteTime: t})
			role.Scripts = append(role.Scripts, &context.ScriptConfig{Name: "公会签到", ScheduleTime: t, CompleteTime: t})
			// role.Scripts = append(role.Scripts, &context.ScriptConfig{Name: "公会聊天", ScheduleTime: t, CompleteTime: t})
			role.Scripts = append(role.Scripts, &context.ScriptConfig{Name: "领取个人奖励", ScheduleTime: t, CompleteTime: t})
			role.Scripts = append(role.Scripts, &context.ScriptConfig{Name: "天空岛贸易", ScheduleTime: t, CompleteTime: t})
			// role.Scripts = append(role.Scripts, &context.ScriptConfig{Name: "武陵道场", ScheduleTime: t, CompleteTime: t})
			role.Scripts = append(role.Scripts, &context.ScriptConfig{Name: "周常副本", ScheduleTime: t, CompleteTime: t})
			// role.Scripts = append(role.Scripts, &context.ScriptConfig{Name: "精英副本", ScheduleTime: t, CompleteTime: t})
			role.Scripts = append(role.Scripts, &context.ScriptConfig{Name: "材料副本", Options: map[string]string{"副本名": "神木村"}, ScheduleTime: t, CompleteTime: t})
			role.Scripts = append(role.Scripts, &context.ScriptConfig{Name: "奈特的金字塔", ScheduleTime: t, CompleteTime: t})
			role.Scripts = append(role.Scripts, &context.ScriptConfig{Name: "领取个人奖励", ScheduleTime: t, CompleteTime: t})
			role.Scripts = append(role.Scripts, &context.ScriptConfig{Name: "领取日常奖励", ScheduleTime: t, CompleteTime: t})
			role.Scripts = append(role.Scripts, &context.ScriptConfig{Name: "领取公会奖励", ScheduleTime: t, CompleteTime: t})
			ctx.Roles = append(ctx.Roles, role)
		}
	}
	ctx.Save()
}

func main() {

	// gen()

	ctx, err := context.Load("context.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, role := range ctx.Roles {
		ctx.SetRole(role)
		// 角色已完成, 则跳过
		if role.FinishTime.Format("2006-01-02") == time.Now().Format("2006-01-02") {
			log.Infof("[角色-%d] 已完成\n", role.Index)
			continue
		}
		scripts.Enter(role.Index)
		ctx.Lanuch()
		for _, script := range role.Scripts {
			ctx.SetScript(script)
			// 脚本已调度，则跳过
			if script.ScheduleTime.Format("2006-01-02") == time.Now().Format("2006-01-02") {
				log.Infof("[角色-%d] [脚本-%s] 已调度\n", role.Index, script.Name)
				continue
			}
			context.ExecuteScript(ctx, script.Name)
		}
		scripts.Exit()
		ctx.Finish()
	}
}
