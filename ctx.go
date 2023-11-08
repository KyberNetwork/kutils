package kutils

import (
	"context"
	"time"
)

type ctxWithoutCancel struct {
	ctx context.Context
}

func (c ctxWithoutCancel) Deadline() (time.Time, bool)       { return time.Time{}, false }
func (c ctxWithoutCancel) Done() <-chan struct{}             { return nil }
func (c ctxWithoutCancel) Err() error                        { return nil }
func (c ctxWithoutCancel) Value(key interface{}) interface{} { return c.ctx.Value(key) }

func CtxWithoutCancel(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}
	return &ctxWithoutCancel{ctx: ctx}
}
