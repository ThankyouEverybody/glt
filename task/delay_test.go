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

func TestDelay1(t *testing.T) {
	println(time.Now().Format(time.DateTime))
	after := time.NewTimer(1 * time.Second)

	cancel, _ := context.WithCancel(context.TODO())
	time.Sleep(2 * time.Second)
	select {
	case t := <-after.C:
		println(t.Format(time.DateTime))
	case <-cancel.Done():
		println(cancel.Err().Error())
	}
	//cancelFunc()

}
