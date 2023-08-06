package task

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewDelayPool(t *testing.T) {
	pool, err := NewDelayPool()
	if err != nil {
		t.Fatalf("new delay pool err: %v", err)
	}
	var wg sync.WaitGroup
	var dt *Delay
	for i := 0; i < 10; i++ {
		ctx := context.WithValue(context.TODO(), "i", i)
		wg.Add(1)
		dt, _ = pool.Put(time.Duration(i)*time.Millisecond+1*time.Second, func(ctx context.Context) {
			fmt.Printf("current time: %s, i: %v\n", time.Now().Format(time.StampNano), ctx.Value("i"))
			wg.Done()
		}, ctx)
	}
	err = pool.Cancel(dt)
	wg.Done()
	time.Sleep(1 * time.Second)
	ctx := context.WithValue(context.TODO(), "i", 1)
	wg.Add(1)
	dt, _ = pool.Put(time.Duration(1)*time.Millisecond+1*time.Second, func(ctx context.Context) {
		fmt.Printf("current time: %s, i: %v\n", time.Now().Format(time.StampNano), ctx.Value("i"))
		wg.Done()
	}, ctx)
	if err != nil {
		t.Errorf("cancel failure: %v", err)
	}

	wg.Wait()
}
