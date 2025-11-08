package ix

import "fmt"

type Rect struct {
	X int64
	Y int64
	W int64
	H int64
}

func (o Rect) String() string {
	return fmt.Sprintf("(%d,%d,%d,%d)", o.X, o.Y, o.W, o.H)
}
