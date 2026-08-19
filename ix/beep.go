package ix

import (
	"fmt"
	"os/exec"
	"sync"
)

var (
	beepEnabledMu sync.RWMutex
	beepEnabled   = false
)

// SetBeepEnabled enables or disables the beep sound.
func SetBeepEnabled(v bool) {
	beepEnabledMu.Lock()
	beepEnabled = v
	beepEnabledMu.Unlock()
}

// BeepEnabled returns whether the beep sound is currently enabled.
func BeepEnabled() bool {
	beepEnabledMu.RLock()
	defer beepEnabledMu.RUnlock()
	return beepEnabled
}

func PlaySystemSound(soundName string) error {
	soundPath := fmt.Sprintf("/System/Library/Sounds/%s.aiff", soundName)
	cmd := exec.Command("afplay", soundPath)
	return cmd.Run()
}

func Beep() {
	if !BeepEnabled() {
		return
	}
	PlaySystemSound("Glass")
}
