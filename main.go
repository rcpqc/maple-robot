package main

import (
	"context"
	"io"
	"os"
	"strconv"
	"time"

	"maple-robot/config"
	"maple-robot/log"
	"maple-robot/scripts"
)

func main() {
	logFile := "logs/" + time.Now().Format(time.DateOnly) + ".log"
	records, _ := config.LoadTaskRecords(logFile)

	ctx := config.WithRecords(context.Background(), records)

	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	baseLogger := log.New(io.MultiWriter(os.Stdout, f))
	cfg, err := config.Load("config.yaml")
	if err != nil {
		panic(err)
	}

	entered := false
	for _, role := range cfg.Roles {

		ctx := log.WithLogger(ctx, baseLogger.With("role", role.Id, "class", role.Class))

		// 角色已完成, 则跳过
		if _, ok := config.GetRecord(ctx, role.Id, "", "角色日常结束"); ok {
			log.Info(ctx, "角色日常跳过")
			continue
		}

		// 脚本加载
		if role.Script == "" {
			role.Script = "script_150.yaml"
		}
		script, err := config.LoadScript(role.Script)
		if err != nil {
			log.Error(ctx, "脚本加载", "script", role.Script, "err", err)
		}

		// 角色登录
		log.Info(ctx, "角色日常开始")
		if !entered {
			index, _ := strconv.Atoi(role.Id)
			scripts.Enter(ctx, index)
			scripts.WaitEnter(ctx)
			entered = true
		} else {
			scripts.NextRole(ctx)
			scripts.WaitEnter(ctx)
		}

		// 任务执行
		for _, task := range script.Tasks {
			if !task.Condition.Match() {
				continue
			}
			// 任务今日已入场
			if _, ok := config.GetRecord(ctx, role.Id, task.Name, "任务完成"); ok {
				continue
			}
			if _, ok := config.GetRecord(ctx, role.Id, task.Name, "任务入场"); ok {
				continue
			}
			task.Execute(ctx)
		}
		log.Info(ctx, "角色日常结束")
	}
	scripts.LabelClick(ctx, "更改角色-选择角色")
}
