package log

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"time"
)

type Handler struct {
	handler slog.Handler
	Output  io.Writer
}

func NewHandler(w io.Writer) *Handler {
	opts := &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey || a.Key == slog.LevelKey || a.Key == slog.MessageKey {
				return slog.Attr{}
			}
			return a
		},
	}
	return &Handler{
		handler: slog.NewTextHandler(w, opts),
		Output:  w,
	}
}

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{handler: h.handler.WithAttrs(attrs), Output: h.Output}
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{handler: h.handler.WithGroup(name), Output: h.Output}
}

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(time.Now().Format(time.DateTime))
	buf.WriteString(" ")
	buf.WriteString(r.Level.String())
	buf.WriteString(" ")
	buf.WriteString(r.Message)
	buf.WriteString(" ")
	h.Output.Write(buf.Bytes())
	return h.handler.Handle(ctx, r)
}
