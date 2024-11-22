package log

import (
	"fmt"
	"time"
)

func Infof(format string, args ...any) {
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf(" [INFO] "+format, args...)
}

func Warnf(format string, args ...any) {
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf(" [WARN] "+format, args...)
	fnMessageBeep.Call(0xffffffff)
}

func Errorf(format string, args ...any) {
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf(" [ERROR] "+format, args...)
	fnMessageBeep.Call(0xffffffff)
}

func Fatalf(format string, args ...any) {
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf(" [FATAL] "+format, args...)
	fnMessageBeep.Call(0xffffffff)
	panic(-1)
}

func Printf(format string, args ...any) {
	fmt.Printf(format, args...)
}
