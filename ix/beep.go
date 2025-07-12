package ix

import "syscall"

var (
	fnMessageBeep = syscall.MustLoadDLL("user32.dll").MustFindProc("MessageBeep")
)

func Beep() {
	fnMessageBeep.Call(0xffffffff)
}
