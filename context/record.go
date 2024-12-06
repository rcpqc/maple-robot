package context

import (
	"fmt"
	"time"

	"gopkg.in/yaml.v3"
)

type Record struct {
	Time time.Time     `yaml:"time"`
	Cost time.Duration `yaml:"cost"`
}

type Records map[string]*Record

func (o Records) Start(name string) {
	r, ok := o[name]
	if !ok {
		r = &Record{}
		o[name] = r
	}
	r.Time = time.Now()
}

func (o Records) Mark(name string) {
	r, ok := o[name]
	if !ok {
		r = &Record{}
		o[name] = r
	}
	r.Cost = time.Since(r.Time)
}

func (o Records) DailyDone(name string) bool {
	r, ok := o[name]
	if !ok {
		return false
	}
	now := time.Now()
	zero := time.Date(now.Year(), now.Month(), now.Hour(), 0, 0, 0, 0, time.Local)
	return r.Time.After(zero)
}

func (o *Record) MarshalYAML() (interface{}, error) {
	if o.Cost == 0 {
		return "", nil
	}
	return fmt.Sprintf("%s + %.0f", o.Time.Format("2006-01-02T15:04:05"), o.Cost.Seconds()), nil
}

func (o *Record) UnmarshalYAML(value *yaml.Node) error {
	line := value.Value
	if line == "" {
		o.Time = time.Unix(0, 0)
		o.Cost = 0
		return nil
	}
	if len(line) < 23 {
		return fmt.Errorf("illegal record = %s", line)
	}
	var t string
	var cost int64
	_, err := fmt.Sscanf(value.Value, "%s + %d", &t, &cost)
	if err != nil {
		return err
	}
	o.Time, _ = time.Parse("2006-01-02T15:04:05", t)
	o.Cost = time.Duration(cost) * time.Second
	return nil
}
