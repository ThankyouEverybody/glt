package task

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestDelay(t *testing.T) {
	var wait sync.WaitGroup
	t.Logf("currentTime: %s", time.Now().String())
	wait.Add(1)
	delay, _ := NewDelay(5*time.Second, func(ctx context.Context) {
		t.Logf("delay start: %s", time.Now().String())
		wait.Done()
	})
	t.Logf("status: %d,currentTime: %s", delay.Status(), time.Now().String())
	delay, _ = NewDelay(5*time.Second, func(ctx context.Context) {
		t.Logf("delay start: %s", time.Now().String())
		wait.Done()
	})
	if err := delay.Cancel(); err != nil {
		t.Logf("first err: %v", err)

	}
	if err := delay.Cancel(); err != nil {
		t.Logf("second err: %v", err)
	}

	wait.Wait()
}
