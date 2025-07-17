package config

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"
)

type recordkey struct{}

type Records map[string]time.Time

var Events = map[string]bool{
	"任务开始":   true,
	"任务完成":   true,
	"任务入场":   true,
	"角色日常开始": true,
	"角色日常结束": true,
}

// WithRecords
func WithRecords(ctx context.Context, records Records) context.Context {
	return context.WithValue(ctx, recordkey{}, records)
}

// GetRecord
func GetRecord(ctx context.Context, role string, task string, event string) (time.Time, bool) {
	records, ok := ctx.Value(recordkey{}).(Records)
	if !ok {
		return time.Time{}, false
	}
	key := role + "-" + task + "-" + event
	t, ok := records[key]
	return t, ok
}

// ParseRecord 解析日志记录
func ParseRecord(line string) (*slog.Record, error) {
	fields := strings.Split(line, " ")
	r := &slog.Record{}
	// 解析时间
	t, err := time.ParseInLocation(time.DateTime, fields[0]+" "+fields[1], time.Local)
	if err != nil {
		return nil, fmt.Errorf("time.parse(%s %s): %w", fields[0], fields[1], err)
	}
	r.Time = t
	// 解析日志等级
	switch fields[2] {
	case "DEBUG":
		r.Level = slog.LevelDebug
	case "INFO":
		r.Level = slog.LevelInfo
	case "WARN":
		r.Level = slog.LevelWarn
	case "ERROR":
		r.Level = slog.LevelError
	default:
		return nil, fmt.Errorf("unknown level=%s", fields[2])
	}
	// 解析消息
	r.Message = fields[3]
	// 解析属性
	for i := 4; i < len(fields); i++ {
		kv := strings.Split(fields[i], "=")
		if len(kv) != 2 {
			return nil, fmt.Errorf("kv.parse(%s) fail", fields[4])
		}
		r.AddAttrs(slog.String(kv[0], kv[1]))
	}
	return r, nil
}

// LoadTaskRecords 加载记录, 角色-任务-事件=时间戳
func LoadTaskRecords(file string) (Records, error) {

	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("open(%s): %w", file, err)
	}
	defer f.Close()

	records := Records{}

	br := bufio.NewReaderSize(f, 0x1000)
	for i := 0; ; i++ {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("readline(%d): %w", i, err)
		}
		r, err := ParseRecord(string(line))
		if err != nil {
			return nil, fmt.Errorf("parse(%s): %w", string(line), err)
		}
		event := r.Message
		if _, ok := Events[event]; !ok {
			continue
		}
		var role, task string
		r.Attrs(func(a slog.Attr) bool {
			switch a.Key {
			case "role":
				role = a.Value.String()
			case "task":
				task = a.Value.String()
			}
			return true
		})
		records[role+"-"+task+"-"+event] = r.Time
	}
	return records, nil
}
