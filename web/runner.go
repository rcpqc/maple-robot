package web

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"maple-robot/config"
	"maple-robot/log"
	"maple-robot/scripts"

	"github.com/rcpqc/expr"
)

// Runner controls the robot execution loop.
type Runner struct {
	mu     sync.Mutex
	cancel context.CancelFunc

	currentRole string
	currentTask string
	roleIdx     int
	totalRoles  int
	running     bool
	startTime   time.Time

	logHub *LogHub
}

// NewRunner creates a Runner. Roles, log file and records are loaded on each
// Start() so that cross-day launches pick up the current day's log file and
// the latest scripts.yaml.
func NewRunner(logHub *LogHub) *Runner {
	return &Runner{
		logHub: logHub,
	}
}

// StatusResp is returned by the status API.
type StatusResp struct {
	Running     bool   `json:"running"`
	CurrentRole string `json:"current_role"`
	CurrentTask string `json:"current_task"`
	RoleIdx     int    `json:"role_idx"`
	TotalRoles  int    `json:"total_roles"`
	StartTime   string `json:"start_time"`
}

// Status returns the current runner state.
func (r *Runner) Status() StatusResp {
	r.mu.Lock()
	defer r.mu.Unlock()
	st := ""
	if !r.startTime.IsZero() {
		st = r.startTime.Format(time.DateTime)
	}
	return StatusResp{
		Running:     r.running,
		CurrentRole: r.currentRole,
		CurrentTask: r.currentTask,
		RoleIdx:     r.roleIdx,
		TotalRoles:  r.totalRoles,
		StartTime:   st,
	}
}

// Start begins executing roles in a background goroutine.
// Returns an error if already running.
func (r *Runner) Start() error {
	r.mu.Lock()
	if r.running {
		r.mu.Unlock()
		return fmt.Errorf("already running")
	}
	ctx, cancel := context.WithCancel(context.Background())
	r.cancel = cancel
	r.running = true
	r.startTime = time.Now()
	r.currentRole = ""
	r.currentTask = ""
	r.roleIdx = 0
	r.totalRoles = 0
	r.mu.Unlock()

	go r.run(ctx)
	return nil
}

// Stop cancels the current execution.
func (r *Runner) Stop() {
	r.mu.Lock()
	cancel := r.cancel
	r.mu.Unlock()
	if cancel != nil {
		cancel()
	}
}

const stopTimeout = 30 * time.Second

// StopWait cancels and waits for execution to finish within the default timeout.
// Returns an error if the runner does not stop within the timeout.
func (r *Runner) StopWait() error {
	return r.StopWaitTimeout(stopTimeout)
}

// StopWaitTimeout cancels execution and waits up to timeout for it to finish.
func (r *Runner) StopWaitTimeout(timeout time.Duration) error {
	r.Stop()
	ddl := time.Now().Add(timeout)
	for {
		r.mu.Lock()
		run := r.running
		r.mu.Unlock()
		if !run {
			return nil
		}
		if time.Now().After(ddl) {
			return fmt.Errorf("stop timeout after %v", timeout)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (r *Runner) run(ctx context.Context) {
	// 按点击当天生成日志文件 (支持跨天启动)
	logFile := "logs/" + time.Now().Format(time.DateOnly) + ".log"
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Printf("[runner] 打开日志文件失败: %v\n", err)
		r.finish()
		return
	}
	defer f.Close()

	baseLogger := log.New(io.MultiWriter(os.Stdout, f, r.logHub))

	// 每次启动重新加载角色 (支持跨天/热更新 scripts.yaml)
	roles, err := config.LoadRoles()
	if err != nil {
		baseLogger.Error("加载角色失败", "err", err)
		r.finish()
		return
	}

	r.mu.Lock()
	r.totalRoles = len(roles)
	r.mu.Unlock()

	// 加载当日任务记录 (按当天日志文件)
	records, _ := config.LoadTaskRecords(logFile)
	ctx = config.WithRecords(ctx, records)

	// 启动时假设游戏已停留在某个已登录角色的世界界面:
	// 第一个需要执行的角色直接使用当前登录角色, 后续角色通过"导航→更改角色"切换。
	first := true

	for i, role := range roles {
		select {
		case <-ctx.Done():
			r.finish()
			return
		default:
		}

		index := (i/7+1)*100 + (i%7 + 1)
		roleid := strconv.FormatInt(int64(index), 10)

		r.mu.Lock()
		r.roleIdx = i + 1
		r.currentRole = role.Class
		r.currentTask = ""
		r.mu.Unlock()

		ctx = scripts.WithRole(ctx, index, role.Class)
		ctx = log.WithLogger(ctx, baseLogger.With("role", index, "class", role.Class))

		// skip if already done today
		if _, ok := config.GetRecord(ctx, roleid, "", "角色日常结束"); ok {
			log.Info(ctx, "角色日常跳过")
			continue
		}

		// load script
		if role.Script == "" {
			role.Script = "200"
		}
		if role.Script == "skip" {
			log.Info(ctx, "角色跳过", "script", "skip")
			continue
		}
		script, err := config.LoadScript(role.Script)
		if err != nil {
			log.Error(ctx, "脚本加载", "script", role.Script, "err", err)
			continue
		}

		// login / switch role
		log.Info(ctx, "角色日常开始")
		if first {
			// 当前登录角色即为首个待执行角色, 直接进入世界
			scripts.WaitEnter(ctx)
			first = false
		} else {
			scripts.NextRole(ctx)
			scripts.WaitEnter(ctx)
		}

		// collect stats
		log.Info(ctx, "角色属性",
			"bpu", fmt.Sprintf("%.0f%%", scripts.BackpackUtilizationRatio()*100),
			"exp", fmt.Sprintf("%.2f%%", scripts.ExpRatio()*100),
		)

		// execute tasks
		for _, task := range script.Tasks {
			select {
			case <-ctx.Done():
				r.finish()
				return
			default:
			}

			r.mu.Lock()
			r.currentTask = task.Key()
			r.mu.Unlock()

			vars := expr.Vars{"index": index, "weekday": int64(time.Now().Weekday())}
			if !task.Condition.Match(vars) {
				continue
			}
			if _, ok := config.GetRecord(ctx, roleid, task.Key(), "任务完成"); ok {
				continue
			}
			if _, ok := config.GetRecord(ctx, roleid, task.Key(), "任务入场"); ok {
				continue
			}
			if config.IsTaskDisabled(role.Script, task.Key()) {
				continue
			}
			task.Execute(ctx)
		}

		r.mu.Lock()
		r.currentTask = ""
		r.mu.Unlock()

		// collect stats after
		log.Info(ctx, "角色属性",
			"bpu", fmt.Sprintf("%.0f%%", scripts.BackpackUtilizationRatio()*100),
			"exp", fmt.Sprintf("%.2f%%", scripts.ExpRatio()*100),
		)
		log.Info(ctx, "角色日常结束")
	}

	// all done
	scripts.Exit(ctx)
	exec.Command("./bin/analyze").Run()
	r.finish()
}

func (r *Runner) finish() {
	r.mu.Lock()
	r.running = false
	r.cancel = nil
	r.mu.Unlock()
}
