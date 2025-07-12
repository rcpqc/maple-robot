package log

import (
	"context"
	"io"
	"log/slog"
	"maple-robot/ix"
)

func New(w io.Writer) *slog.Logger {
	return slog.New(NewHandler(w))
}

func Debug(ctx context.Context, msg string, args ...any) {
	GetLogger(ctx).Debug(msg, args...)
}
func Info(ctx context.Context, msg string, args ...any) {
	GetLogger(ctx).Info(msg, args...)
}

func Warn(ctx context.Context, msg string, args ...any) {
	GetLogger(ctx).Warn(msg, args...)
	ix.Beep()
}

func Error(ctx context.Context, msg string, args ...any) {
	GetLogger(ctx).Error(msg, args...)
	ix.Beep()
}
