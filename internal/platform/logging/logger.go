package logging

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
)

type Options struct {
	ConsoleLevel slog.Level
	FileLevel    slog.Level
	FilePath     string
	AppName      string
}

func Setup(opts Options) (*slog.Logger, error) {
	consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: opts.ConsoleLevel})

	if err := os.MkdirAll(filepath.Dir(opts.FilePath), 0o755); err != nil {
		return nil, err
	}
	file, err := os.OpenFile(opts.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, err
	}
	fileHandler := slog.NewTextHandler(file, &slog.HandlerOptions{Level: opts.FileLevel})

	handler := &teeHandler{handlers: []slog.Handler{consoleHandler, fileHandler}}
	return slog.New(handler).With("app", opts.AppName), nil
}

type teeHandler struct {
	handlers []slog.Handler
}

func (t *teeHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range t.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (t *teeHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, h := range t.handlers {
		if !h.Enabled(ctx, r.Level) {
			continue
		}
		if err := h.Handle(ctx, r); err != nil {
			return err
		}
	}
	return nil
}

func (t *teeHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	hs := make([]slog.Handler, len(t.handlers))
	for i, h := range t.handlers {
		hs[i] = h.WithAttrs(attrs)
	}
	return &teeHandler{handlers: hs}
}

func (t *teeHandler) WithGroup(name string) slog.Handler {
	hs := make([]slog.Handler, len(t.handlers))
	for i, h := range t.handlers {
		hs[i] = h.WithGroup(name)
	}
	return &teeHandler{handlers: hs}
}
