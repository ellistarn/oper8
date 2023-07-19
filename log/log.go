package log

import (
	"context"
	"testing"

	"github.com/samber/lo"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

var key = struct{}{}

func New(ctx context.Context) context.Context {
	return With(ctx, lo.Must(zap.NewProduction()).Sugar())
}

func NewForT(ctx context.Context, t *testing.T) context.Context {
	return With(ctx, zaptest.NewLogger(t).Sugar())
}

func With(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, key, logger)
}

func From(ctx context.Context) *zap.SugaredLogger {
	return ctx.Value(key).(*zap.SugaredLogger)
}
