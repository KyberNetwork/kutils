package kutils

import (
	"context"
	"math"
	"runtime"
	"runtime/debug"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"

	"github.com/KyberNetwork/kutils/klog"
)

//go:generate mockgen -source=batcher.go -destination mocks/mocks.go -package mocks

var (
	ErrBatcherClosed = errors.New("batcher closed")
)

// BatchableTask represents a batchable task
type BatchableTask[R any] interface {
	Ctx() context.Context     // The context of this task
	Done() <-chan struct{}    // Signals if this task was already resolved
	IsDone() bool             // Checks (non-blocking) if this task was already resolved
	Result() (R, error)       // Blocks until this task is resolved and returns result and error
	Resolve(ret R, err error) // Resolves this task with return value and error
}

// ChanTask uses a done channel to signal resolution of return value and error
type ChanTask[R any] struct {
	ctx  context.Context
	done chan struct{}
	Ret  R
	Err  error
}

func NewChanTask[R any](ctx context.Context) *ChanTask[R] {
	if ctx == nil {
		ctx = context.Background()
	}
	return &ChanTask[R]{
		ctx:  ctx,
		done: make(chan struct{}),
	}
}

func (c *ChanTask[R]) Ctx() context.Context {
	return c.ctx
}

func (c *ChanTask[R]) Done() <-chan struct{} {
	return c.done
}

func (c *ChanTask[R]) IsDone() bool {
	select {
	case <-c.done:
		return true
	default:
		return false
	}
}

func (c *ChanTask[R]) Result() (R, error) {
	if c.IsDone() {
		return c.Ret, c.Err
	}
	select {
	case <-c.done:
		return c.Ret, c.Err
	case <-c.ctx.Done():
		return *new(R), c.ctx.Err()
	}
}

func (c *ChanTask[R]) Resolve(ret R, err error) {
	select {
	case <-c.done:
		klog.Errorf(c.ctx, "ChanTask.Resolve|called twice, ignored|c.Ret=%v,c.Err=%v|Ret=%v,Err=%v",
			c.Ret, c.Err, ret, err)
	default:
		c.Ret, c.Err = ret, err
		close(c.done)
	}
}

// Batcher batches together n BatchableTask's together and executes a logic for a batch of BatchableTask's.
// It skips BatchableTask's with cancelled Ctx and resolve those tasks with the context's error.
// Batch logic execution should signal each BatchableTask as done by using its Resolve method.
type Batcher[T BatchableTask[R], R any] interface {
	// Batch submits a BatchableTask to the batcher.
	Batch(task T)
	// Close should stop Batch from being called and clean up any background resources.
	Close()
}

// BatchCfg provides batchRate and batchCnt configs for a ChanBatcher. ChanBatcher will trigger a batch processing
// either if no more task is queued after batchRate, or batchCnt BatchableTask's are already queued.
type BatchCfg func() (batchRate time.Duration, batchCnt int)

// BatchFn is called for a batch of tasks collected and triggered by a ChanBatcher per its batchCfg.
type BatchFn[T any] func([]T)

// ChanBatcher implements Batcher using golang channel.
type ChanBatcher[T BatchableTask[R], R any] struct {
	batchCfg BatchCfg
	batchFn  BatchFn[T]
	taskCh   chan T
	closed   atomic.Bool
}

func NewChanBatcher[T BatchableTask[R], R any](batchCfg BatchCfg, batchFn BatchFn[T]) *ChanBatcher[T, R] {
	_, batchCnt := batchCfg()
	chanBatcher := &ChanBatcher[T, R]{
		batchCfg: batchCfg,
		batchFn:  batchFn,
		taskCh:   make(chan T, 16*batchCnt),
	}
	go chanBatcher.worker()
	return chanBatcher
}

// Batch submits a BatchableTask to the channel if this chanBatcher hasn't been closed.
func (b *ChanBatcher[T, R]) Batch(task T) {
	if !b.closed.Load() {
		b.taskCh <- task
	} else {
		task.Resolve(*new(R), ErrBatcherClosed)
	}
}

// Close closes this chanBatcher to prevents Batch-ing new BatchableTask's and tell the worker goroutine to finish up.
func (b *ChanBatcher[_, _]) Close() {
	if !b.closed.Swap(true) {
		close(b.taskCh)
	}
}

// goBatchFn
func (b *ChanBatcher[T, R]) batchFnWithRecover(tasks []T) {
	defer func() {
		p := recover()
		if p == nil {
			return
		}
		klog.Errorf(context.Background(), "ChanBatcher.goBatchFn|recovered from panic: %v\n%s",
			p, string(debug.Stack()))
		var ret R
		for _, task := range tasks {
			if task.IsDone() {
				continue
			}
			if err, ok := p.(error); ok {
				task.Resolve(ret, errors.Wrap(err, "batchFn panicked"))
			} else {
				task.Resolve(ret, errors.Errorf("batchFn panicked: %v", p))
			}
		}
	}()
	b.batchFn(tasks)
}

// worker batches up BatchableTask's in taskCh per batchCfg (per at most batchRate ns and at most batchCnt BatchableTask's)
// and triggers batchFn with each batch.
func (b *ChanBatcher[T, R]) worker() {
	defer func() {
		if p := recover(); p != nil {
			klog.Errorf(context.Background(), "ChanBatcher.worker|recovered from panic: %v\n%s",
				p, string(debug.Stack()))
		}
	}()
	var tasks []T
	batchTimer := time.NewTimer(time.Duration(math.MaxInt64))
	for {
		runtime.Gosched() // in case GOMAXPROCS is 1, we need to cooperatively yield
		select {
		case <-batchTimer.C:
			if len(tasks) == 0 {
				break
			}
			klog.Debugf(tasks[0].Ctx(), "ChanBatcher.worker|timer|%d tasks", len(tasks))
			go b.batchFnWithRecover(tasks)
			tasks = tasks[:0:0]
		case task, ok := <-b.taskCh:
			ctx := task.Ctx()
			if !ok {
				klog.Debugf(ctx, "ChanBatcher.worker|closed|%d tasks", len(tasks))
				if len(tasks) > 0 {
					go b.batchFnWithRecover(tasks)
				}
				return
			}
			if !task.IsDone() {
				select {
				case <-ctx.Done():
					klog.Infof(ctx, "ChanBatcher.worker|skip|task=%v", task)
					task.Resolve(*new(R), ctx.Err())
					continue
				default:
				}
			}
			duration, batchCount := b.batchCfg()
			if len(tasks) == 0 {
				klog.Debugf(ctx, "ChanBatcher.worker|timer start|duration=%s", duration)
				if !batchTimer.Stop() {
					select {
					case <-batchTimer.C:
					default:
					}
				}
				batchTimer.Reset(duration)
			}
			tasks = append(tasks, task)
			if len(tasks) >= batchCount {
				klog.Debugf(ctx, "ChanBatcher.worker|max|%d tasks", len(tasks))
				go b.batchFnWithRecover(tasks)
				tasks = tasks[:0:0]
			}
		}
	}
}
