package ix

import (
	"context"
	"os"
	"time"
)

var bypass = make(chan time.Time)

func init() {
	go func() {
		for {
			in := []byte{0}
			os.Stdin.Read(in)
			if in[0] == '\n' {
				bypass <- time.Now()
			}
		}
	}()
}

// Bypass sends a bypass signal, causing the current WaitOrPass call to return true.
func Bypass() {
	select {
	case bypass <- time.Now():
	default:
	}
}

// WaitOrPass waits up to timeout for either a bypass signal or context cancellation.
// Returns true if bypassed or cancelled, false on timeout.
func WaitOrPass(ctx context.Context, timeout time.Duration) bool {
	now := time.Now()
	for {
		select {
		case <-ctx.Done():
			return true
		case t := <-bypass:
			if t.After(now) {
				return true
			} else {
				continue
			}
		case <-time.After(timeout):
			return false
		}
	}
}
