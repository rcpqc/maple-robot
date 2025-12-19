package ix

import (
	"fmt"
	"os/exec"
)

func PlaySystemSound(soundName string) error {
	soundPath := fmt.Sprintf("/System/Library/Sounds/%s.aiff", soundName)
	cmd := exec.Command("afplay", soundPath)
	return cmd.Run()
}

func Beep() {
	PlaySystemSound("Glass")
}
