package log

import "syscall"

var (
	fnMessageBeep = syscall.MustLoadDLL("user32.dll").MustFindProc("MessageBeep")
)
