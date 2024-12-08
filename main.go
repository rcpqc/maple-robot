package main

import (
	"maple-robot/context"
	"maple-robot/log"
	"maple-robot/scripts"
)

func main() {
	ctx, err := context.Load("context.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, role := range ctx.Roles {
		if role.Records == nil {
			role.Records = make(context.Records)
		}
		ctx.SetRole(role)
		// 角色已完成, 则跳过
		if role.Records.DailyDone("角色") {
			log.Infof("[角色-%d] 已完成\n", role.Index)
			continue
		}
		script, err := scripts.Load(role.Script)
		if err != nil {
			log.Fatalf("[角色-%d] 加载脚本(%s) -> %v", role.Index, role.Script, err)
		}
		scripts.Enter(role.Index)
		ctx.Lanuch()
		ctx.Execute(script.Tasks)
		scripts.Exit()
		ctx.Finish()
	}
}
