package kutils

import (
	"context"
	"go.uber.org/zap"
	"sync"
	"testing"

	"github.com/KyberNetwork/kutils/klog"
)

func TestZapLoggerDataRace(t *testing.T) {
	ctx := context.Background()
	//logger := LoggerFromCtx(ctx)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			klog.Infof(ctx, "Logging in goroutine %v", zap.Int("goroutine", i))
		}(i)
	}
	wg.Wait()
}
