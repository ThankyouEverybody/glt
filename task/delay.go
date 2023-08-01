package task

import (
	"context"
	"errors"
	"fmt"
	"github.com/cafe-old-babe/glt/base"
	"strconv"
	"sync/atomic"
	"time"
)

type Delay struct {
	ctx      context.Context
	do       base.CtxFunc
	cancel   context.CancelFunc
	duration time.Duration
	status   atomic.Int64 // 0 wait ,1,running,2:after running done,3:after running err 4:cancel
}

func NewDelay(duration time.Duration, delayFunc base.CtxFunc,
	ctx ...context.Context) (task *Delay, err error) {

	if duration <= 0 {
		err = errors.New("arg is invalid")
		return
	}
	task = new(Delay)

	if ctx == nil || len(ctx) <= 0 {
		task.ctx = context.Background()
	} else {
		task.ctx = ctx[0]
	}
	task.ctx, task.cancel = context.WithCancel(task.ctx)
	task.do = delayFunc
	task.duration = duration
	task.status.Store(Wait)
	go func(task *Delay) {
		defer func() {
			if err := recover(); err != nil {
				task.status.CompareAndSwap(Running, AfterRunningErr)
				return
			}
			if !task.status.CompareAndSwap(Running, AfterRunningDone) {
				task.status.Store(Cancel)
			}

		}()
		select {
		case <-time.After(task.duration):
			if task.status.CompareAndSwap(Wait, Running) {
				task.do(task.ctx)
			}
		case <-task.ctx.Done():
			return
		}

	}(task)
	return
}

func (_self *Delay) Cancel() (err error) {
	if _self.status.CompareAndSwap(Wait, Cancel) {
		_self.cancel()
		return
	}
	format := "task cancel failure, current status: %s"
	switch _self.status.Load() {
	case Running:
		err = errors.New(fmt.Sprintf(format, "running"))
	case AfterRunningDone:
		err = errors.New(fmt.Sprintf(format, "after runing done"))
	case AfterRunningErr:
		err = errors.New(fmt.Sprintf(format, "after running error"))
	case Cancel:
		err = errors.New(fmt.Sprintf(format, "cancel"))

	default:
		err = errors.New("unknown Error, status: " + strconv.FormatInt(_self.status.Load(), 10))
	}
	return
}

func (_self *Delay) Status() int {
	return int(_self.status.Load())
}
