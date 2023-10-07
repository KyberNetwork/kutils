package kutils

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCtxWithoutCancel(t *testing.T) {
	cancelledCtx, cancel := context.WithTimeout(
		context.WithValue(CtxWithoutCancel(nil), "key", "value"), time.Second)
	detachedCtx := CtxWithoutCancel(cancelledCtx)

	t.Run("Deadline()", func(t *testing.T) {
		_, ok := cancelledCtx.Deadline()
		assert.True(t, ok, "cancelledCtx should have Deadline")
		_, ok = detachedCtx.Deadline()
		assert.False(t, ok, "detachedCtx should not have Deadline")
		cancel()
	})

	t.Run("Err()", func(t *testing.T) {
		assert.NotNil(t, cancelledCtx.Err())
		assert.Nil(t, detachedCtx.Err())
	})

	t.Run("Done()", func(t *testing.T) {
		select {
		case <-cancelledCtx.Done():
		default:
			assert.Fail(t, "cancelledCtx.Done() should be closed")
		}
		select {
		case <-detachedCtx.Done():
			assert.Fail(t, "detachedCtx.Done() should not be closed")
		default:
		}
	})

	t.Run("Value()", func(t *testing.T) {
		assert.Equal(t, "value", cancelledCtx.Value("key"))
		assert.Equal(t, "value", detachedCtx.Value("key"))
	})
}
