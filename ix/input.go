package ix

import (
	"os/exec"
	"strconv"
)

type KeyCode int64

const (
	KeyCodeEscape KeyCode = 111
)

func Key(code KeyCode) {
	exec.Command("adb", "shell", "input", "keyevent", strconv.FormatInt(int64(code), 10)).Run()
}

func Tap(pos Position) {
	exec.Command("adb", "shell", "input", "tap", strconv.FormatInt(pos.X, 10), strconv.FormatInt(pos.Y, 10)).Run()
}

func Text(text string) {
	exec.Command("adb", "shell", "input", "text", text).Run()
}

func Swipe(posa Position, posb Position, durationMs int64) {
	exec.Command("adb", "shell", "input", "swipe",
		strconv.FormatInt(posa.X, 10), strconv.FormatInt(posa.Y, 10),
		strconv.FormatInt(posb.X, 10), strconv.FormatInt(posb.Y, 10),
		strconv.FormatInt(durationMs, 10),
	).Run()
}
