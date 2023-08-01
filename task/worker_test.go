package task

import (
	"context"
	"errors"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	var success atomic.Int32
	var failure atomic.Int32
	var wait sync.WaitGroup
	worker, err := NewNumWorker("1", func(ctx context.Context, err error) {
		value := ctx.Value("errVal")
		failure.Add(1)
		t.Logf("errVal: %v", value)
	}, 10)
	if err != nil {
		t.Fatalf("new worker err: %v", err)
	}

	for i := 0; i < 100; i++ {
		wait.Add(1)
		go func(j int, wg *sync.WaitGroup) {
			worker.Do(func() (context.Context, error) {
				defer wg.Done()
				if j%10 == 0 {
					ctx := context.WithValue(context.TODO(), "errVal", "errVal-"+strconv.Itoa(j))
					return ctx, errors.New("errNo: %s" + strconv.Itoa(j))
				}
				now := time.Now()
				for time.Now().Before(now.Add(100 * time.Millisecond)) {

				}

				success.Add(1)
				return nil, nil
			})
		}(i, &wait)
	}
	wait.Wait()
	err = worker.DoOver()
	if err != nil {
		t.Errorf("do over error: %v", err)
	}
	worker.DoWait()
	t.Logf("sucess: %d, fialure: %d", success.Load(), failure.Load())
}
