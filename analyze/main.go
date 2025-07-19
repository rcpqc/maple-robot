package main

import (
	"log"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"maple-robot/config"
)

type Interval struct {
	StartTime time.Time
	EndTime   time.Time
}

func (o Interval) Seconds() int {
	if o.EndTime.IsZero() {
		return 0
	}
	return int(o.EndTime.Sub(o.StartTime).Seconds())
}

type Task struct {
	Name     string
	Interval Interval
}

type Role struct {
	Id        string
	Class     string
	Intervals []Interval
	Tasks     map[string]*Task
}

func main() {
	date := "2025-07-19"
	records, err := config.LoadRecords("../logs/" + date + ".log")
	if err != nil {
		log.Fatal(err)
	}

	// 按日志时间顺序，统计角色各任务阶段耗时
	mroles := map[string]*Role{}
	for _, r := range records {
		id := config.GetRecordAttr(r, "role")
		switch r.Message {
		case "角色日常开始":
			// 已经存在，说明该角色是中断后继续，直接增加时间间隔
			if role, ok := mroles[id]; ok {
				role.Intervals = append(role.Intervals, Interval{StartTime: r.Time})
			} else {
				class := config.GetRecordAttr(r, "class")
				mroles[id] = &Role{
					Id:        id,
					Class:     class,
					Intervals: []Interval{{StartTime: r.Time}},
					Tasks:     map[string]*Task{},
				}
			}
		case "任务开始":
			task := config.GetRecordAttr(r, "task")
			role := mroles[id]
			role.Tasks[task] = &Task{Name: task, Interval: Interval{StartTime: r.Time}}
		case "任务完成":
			task := config.GetRecordAttr(r, "task")
			role := mroles[id]
			role.Tasks[task].Interval.EndTime = r.Time
			role.Intervals[len(role.Intervals)-1].EndTime = r.Time
		case "角色日常结束":
			role := mroles[id]
			role.Intervals[len(role.Intervals)-1].EndTime = r.Time
		}
	}

	roles := slices.Collect(maps.Values(mroles))
	slices.SortFunc(roles, func(a, b *Role) int { return strings.Compare(a.Id, b.Id) })

	mnames := map[string]struct{}{}
	for _, role := range roles {
		for _, task := range role.Tasks {
			mnames[task.Name] = struct{}{}
		}
	}
	names := slices.Collect(maps.Keys(mnames))
	slices.Sort(names)

	// 输出
	fOut, _ := os.Create("./reports/" + date + ".csv")
	headers := []string{"id", "角色", "角色耗时(s)", "任务耗时(s)", "角色切换"}
	headers = append(headers, names...)
	fOut.WriteString(strings.Join(headers, ",") + "\n")
	for _, role := range roles {
		// 计算角色总耗时
		roleCost := 0
		for _, interval := range role.Intervals {
			roleCost += interval.Seconds()
		}
		// 计算各个任务耗时
		taskCost := 0
		columns := []string{role.Id, role.Class, "", "", ""}
		for _, name := range names {
			if task, ok := role.Tasks[name]; ok {
				cost := task.Interval.Seconds()
				taskCost += cost
				columns = append(columns, strconv.Itoa(int(cost)))
			} else {
				columns = append(columns, "0")
			}
		}
		columns[2] = strconv.Itoa(roleCost)
		columns[3] = strconv.Itoa(taskCost)
		columns[4] = strconv.Itoa(roleCost - taskCost)
		fOut.WriteString(strings.Join(columns, ",") + "\n")
	}
	fOut.Close()
}
