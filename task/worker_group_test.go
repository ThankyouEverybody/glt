package task

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
)

func TestWorkerGroup(t *testing.T) {
	var success, failure atomic.Int32
	ctx := context.TODO()
	group, err := NewWorkerGroup(ctx, 10, func(ctx context.Context) (string, error) {
		value := ctx.Value("key")
		return value.(string), nil
	}, func(ctx context.Context, err error) {
		failure.Add(1)
	})
	if err != nil {
		t.Fatalf("new work group err: %v", err)
	}
	var waitGroup sync.WaitGroup
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go func(j int, wg *sync.WaitGroup) {
			withValue := context.WithValue(ctx, "key", strconv.Itoa(j%10))
			defer wg.Done()
			_ = group.Push(withValue, func() (context.Context, error) {
				value := withValue.Value("key")
				if value.(string) == "0" {
					return nil, errors.New("test err: 0")
				}
				success.Add(1)
				return nil, nil
			})
		}(i, &waitGroup)

	}
	waitGroup.Wait()
	group.Wait()
	fmt.Printf("success: %d, failure: %d\n", success.Load(), failure.Load())
}
