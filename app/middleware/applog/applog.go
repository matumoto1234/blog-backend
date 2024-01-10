package applog

import (
	"context"
	"log/slog"
	"os"
)

func New() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

type contextKey struct{}

func WithContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

func FromContext(ctx context.Context) *slog.Logger {
	l, ok := ctx.Value(contextKey{}).(*slog.Logger)
	if !ok {
		panic("logger not found in context")
	}
	return l
}
