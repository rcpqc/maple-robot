package log

import (
	"context"
	"log/slog"
)

type ctxkey struct{}

func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxkey{}, logger)
}

func GetLogger(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(ctxkey{}).(*slog.Logger)
	if ok {
		return logger
	}
	return slog.Default()
}
