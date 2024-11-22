package ix

import "fmt"

type Position struct {
	X int64
	Y int64
}

func (o Position) String() string {
	return fmt.Sprintf("(%d,%d)", o.X, o.Y)
}
