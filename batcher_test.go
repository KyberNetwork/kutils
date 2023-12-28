package kutils

import (
	"context"
	"runtime"
	"sync/atomic"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/KyberNetwork/kutils/klog"
)

func TestChanBatcher(t *testing.T) {
	ctx := context.Background()
	batchRate := 10 * time.Millisecond
	batchFn := func(_ []*ChanTask[time.Duration]) {}
	batcher := NewChanBatcher[*ChanTask[time.Duration], time.Duration](func() (time.Duration, int) {
		return batchRate, 2
	}, func(tasks []*ChanTask[time.Duration]) { batchFn(tasks) })
	var cnt atomic.Uint32
	start := time.Now()
	batchFn = func(tasks []*ChanTask[time.Duration]) {
		cnt.Add(1)
		for _, task := range tasks {
			task.Resolve(time.Since(start), nil)
		}
	}
	task0 := NewChanTask[time.Duration](ctx)
	task1 := NewChanTask[time.Duration](ctx)
	task2 := NewChanTask[time.Duration](ctx)

	t.Run("happy", func(t *testing.T) {
		batcher.Batch(task0)
		batcher.Batch(task1)
		_, _ = task0.Result()
		assert.EqualValues(t, 1, cnt.Load())
		assert.NoError(t, task0.Err)
		assert.Less(t, task0.Ret, batchRate)
		ret, err := task1.Result()
		assert.NoError(t, err)
		assert.Less(t, ret, batchRate)
		time.Sleep(batchRate * 11 / 10)
		runtime.Gosched()

		batcher.Batch(task2)
		assert.False(t, task2.IsDone())
		ret, err = task2.Result()
		assert.True(t, task2.IsDone())
		assert.EqualValues(t, 2, cnt.Load())
		assert.Equal(t, task2.Err, err)
		assert.NoError(t, task2.Err)
		assert.Equal(t, task2.Ret, ret)
		assert.Greater(t, ret, batchRate)
	})

	t.Run("spam", func(t *testing.T) {
		batcher := NewChanBatcher[*ChanTask[int], int](func() (time.Duration, int) { return 0, 0 },
			func(tasks []*ChanTask[int]) {
				for _, task := range tasks {
					task.Resolve(0, nil)
				}
			})
		const taskCnt = 1000
		tasks := make([]*ChanTask[int], taskCnt)
		start := time.Now()
		for i := 0; i < taskCnt; i++ {
			tasks[i] = NewChanTask[int](ctx)
			batcher.Batch(tasks[i])
		}
		// 1k: 2.561804ms; 1M: 2.62s - average overhead per task = 2.6µs
		klog.Warnf(ctx, "done %d tasks in %v", taskCnt, time.Since(start))
		for i := 0; i < taskCnt; i++ {
			ret, err := tasks[i].Result()
			assert.NoError(t, err)
			assert.EqualValues(t, 0, ret)
		}
		batcher.Close()
	})

	t.Run("resolve twice", func(t *testing.T) {
		task0.Resolve(batchRate, nil)
		assert.NoError(t, task0.Err)
		assert.Less(t, task0.Ret, batchRate)
	})

	t.Run("recover from panic", func(t *testing.T) {
		oldBatchFn := batchFn
		batchFn = func(tasks []*ChanTask[time.Duration]) {
			panic("test panic")
		}
		task0 = NewChanTask[time.Duration](ctx)
		task1 = NewChanTask[time.Duration](ctx)
		task0.Resolve(0, nil)
		batcher.Batch(task0)
		batcher.Batch(task1)
		<-task1.Done()
		assert.ErrorContains(t, task1.Err, "test panic")

		panicErr := errors.New("test panic error")
		batchFn = func(tasks []*ChanTask[time.Duration]) {
			panic(panicErr)
		}
		task0 = NewChanTask[time.Duration](ctx)
		task1 = NewChanTask[time.Duration](ctx)
		batcher.Batch(task0)
		batcher.Batch(task1)
		<-task1.Done()
		assert.ErrorIs(t, task0.Err, panicErr)
		assert.ErrorIs(t, task1.Err, panicErr)

		batchFn = oldBatchFn
		task2 = NewChanTask[time.Duration](nil) // nolint:staticcheck
		batcher.Batch(task2)
		batcher.Batch(task2)
		ret, err := task2.Result()
		assert.NoError(t, err)
		assert.Greater(t, ret, batchRate)
	})

	t.Run("cancelled task", func(t *testing.T) {
		ctx, cancel := context.WithCancel(ctx)
		task0 = NewChanTask[time.Duration](ctx)
		batcher.Batch(task0)
		cancel()
		_, err := task0.Result()
		assert.ErrorIs(t, err, context.Canceled)
	})

	t.Run("close", func(t *testing.T) {
		batcher.Batch(task2)
		batcher.Close()
		task3 := NewChanTask[time.Duration](ctx)
		batcher.Batch(task3)
		assert.ErrorIs(t, task3.Err, ErrBatcherClosed)
	})

	t.Run("invalid task", func(t *testing.T) {
		NewChanBatcher[*ChanTask[int], int](func() (time.Duration, int) { return 0, 0 },
			nil).Batch(&ChanTask[int]{})
	})
}
