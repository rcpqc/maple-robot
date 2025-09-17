package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"time"

	"maple-robot/config"
	"maple-robot/ix"
	"maple-robot/log"
	"maple-robot/scripts"
)

func init() {
	if err := ix.Check(); err != nil {
		panic(err)
	}
}

func main() {
	// 加载日志记录
	logFile := "logs/" + time.Now().Format(time.DateOnly) + ".log"
	records, _ := config.LoadTaskRecords(logFile)
	ctx := config.WithRecords(context.Background(), records)

	// 设置日志输出
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	baseLogger := log.New(io.MultiWriter(os.Stdout, f))

	// 读取角色脚本
	cfg, err := config.Load("config.yaml")
	if err != nil {
		panic(err)
	}

	entered := false

	for i, role := range cfg.Roles {
		index := (i/7+1)*100 + (i%7 + 1)
		roleid := strconv.FormatInt(int64(index), 10)
		ctx = scripts.WithRole(ctx, index, role.Class)
		ctx = log.WithLogger(ctx, baseLogger.With("role", index, "class", role.Class))
		// 角色已完成, 则跳过
		if _, ok := config.GetRecord(ctx, roleid, "", "角色日常结束"); ok {
			log.Info(ctx, "角色日常跳过")
			continue
		}

		// 脚本加载
		if role.Script == "" {
			role.Script = "script_180.yaml"
		}
		script, err := config.LoadScript(role.Script)
		if err != nil {
			log.Error(ctx, "脚本加载", "script", role.Script, "err", err)
		}

		// 角色登录
		log.Info(ctx, "角色日常开始")
		if !entered {
			scripts.Enter(ctx, index)
			scripts.WaitEnter(ctx)
			entered = true
		} else {
			scripts.NextRole(ctx)
			scripts.WaitEnter(ctx)
		}

		// 角色属性获取
		sExp := scripts.ExpRatio()

		// 任务执行
		for _, task := range script.Tasks {
			if !task.Condition.Match() {
				continue
			}
			// 任务今日已入场
			if _, ok := config.GetRecord(ctx, roleid, task.Name, "任务完成"); ok {
				continue
			}
			if _, ok := config.GetRecord(ctx, roleid, task.Name, "任务入场"); ok {
				continue
			}
			task.Execute(ctx)
		}
		eExp := scripts.ExpRatio()
		if eExp < sExp {
			eExp = 1 + eExp
		}

		log.Info(ctx, "角色属性",
			"bpu", fmt.Sprintf("%.0f%%", scripts.BackpackUtilizationRatio()*100),
			"exp", fmt.Sprintf("%.2f%%", eExp*100),
			"dexp", fmt.Sprintf("%.2f%%", (eExp-sExp)*100),
		)

		log.Info(ctx, "角色日常结束")
	}
	scripts.Exit(ctx)

	f.Close()

	exec.Command("./analyze.exe").Run()

}
