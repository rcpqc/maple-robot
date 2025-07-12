package main

import (
	"context"
	"io"
	"os"
	"time"

	"maple-robot/log"
	"maple-robot/record"
	"maple-robot/scripts"
)

func main() {
	f, err := os.OpenFile("logs/"+time.Now().Format(time.DateOnly)+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	baseLogger := log.New(io.MultiWriter(os.Stdout, f))

	c, err := record.Load("context.yaml")
	if err != nil {
		panic(err)
	}
	entered := false
	for _, role := range c.Roles {
		ctx := log.WithLogger(context.Background(), baseLogger.With("role", role.Index))
		if role.Records == nil {
			role.Records = make(record.Records)
		}
		c.SetRole(role)
		// 角色已完成, 则跳过
		if role.Records.DailyDone("角色") {
			log.Info(ctx, "角色日常完成")
			continue
		}
		log.Info(ctx, "角色日常开始")
		script, err := scripts.Load(role.Script)
		if err != nil {
			log.Error(ctx, "脚本加载", "script", role.Script, "err", err)
		}
		if !entered {
			scripts.Enter(ctx, role.Index)
			entered = true
		} else {
			scripts.WaitEnter(ctx)
		}
		c.Lanuch()
		c.Execute(ctx, script.Tasks)
		scripts.NextRole(ctx)
		log.Info(ctx, "角色日常结束")
		c.Finish()
	}
}
