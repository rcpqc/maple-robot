package ix

import (
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

func WaitOrPass(timeout time.Duration) bool {
	now := time.Now()
	for {
		select {
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
