package scripts

import (
	"context"
)

type ctxkey struct{}

type Message struct {
	Index int
	Class string
}

func WithRole(ctx context.Context, index int, class string) context.Context {
	return context.WithValue(ctx, ctxkey{}, &Message{index, class})
}

func GetRole(ctx context.Context) *Message {
	msg, ok := ctx.Value(ctxkey{}).(*Message)
	if ok {
		return msg
	}
	return nil
}
