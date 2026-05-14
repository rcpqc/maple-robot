package web

import (
	"context"
	"fmt"
	"log/slog"
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

	logFilePath string
	baseLogger  *slog.Logger
	cfg         *config.Config
}

// NewRunner creates a Runner that will execute roles from cfg.
func NewRunner(cfg *config.Config, baseLogger *slog.Logger, logFilePath string) *Runner {
	return &Runner{
		cfg:         cfg,
		baseLogger:  baseLogger,
		logFilePath: logFilePath,
		totalRoles:  len(cfg.Roles),
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

// StopWait cancels and waits for execution to finish.
func (r *Runner) StopWait() {
	r.Stop()
	for {
		r.mu.Lock()
		run := r.running
		r.mu.Unlock()
		if !run {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (r *Runner) run(ctx context.Context) {
	records, _ := config.LoadTaskRecords(r.logFilePath)
	ctx = config.WithRecords(ctx, records)

	entered := false

	for i, role := range r.cfg.Roles {
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
		ctx = log.WithLogger(ctx, r.baseLogger.With("role", index, "class", role.Class))

		// skip if already done today
		if _, ok := config.GetRecord(ctx, roleid, "", "角色日常结束"); ok {
			log.Info(ctx, "角色日常跳过")
			continue
		}

		// load script
		if role.Script == "" {
			role.Script = "script_200.yaml"
		}
		script, err := config.LoadScript(role.Script)
		if err != nil {
			log.Error(ctx, "脚本加载", "script", role.Script, "err", err)
			continue
		}

		// login / switch role
		log.Info(ctx, "角色日常开始")
		if !entered {
			scripts.Enter(ctx, index)
			scripts.WaitEnter(ctx)
			entered = true
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
			r.currentTask = task.Name
			r.mu.Unlock()

			vars := expr.Vars{"index": index, "weekday": int64(time.Now().Weekday())}
			if !task.Condition.Match(vars) {
				continue
			}
			if _, ok := config.GetRecord(ctx, roleid, task.Name, "任务完成"); ok {
				continue
			}
			if _, ok := config.GetRecord(ctx, roleid, task.Name, "任务入场"); ok {
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
