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
	status   *atomic.Int64 // 0 wait ,1,running,2:after running done,3:after running err 4:cancel
}

func newDelayPtr(duration time.Duration, delayFunc base.CtxFunc,
	ctx ...context.Context) (task *Delay, err error) {

	if duration <= 0 {
		err = errors.New("duration then less and equals zero")
		return
	}
	if delayFunc == nil {
		err = errors.New("do func is nil")
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
	var status atomic.Int64
	status.Store(Wait)
	task.status = &status
	return
}

func NewDelay(duration time.Duration, delayFunc base.CtxFunc,
	ctx ...context.Context) (task *Delay, err error) {
	task, err = newDelayPtr(duration, delayFunc, ctx...)
	if err != nil {
		return
	}
	go task.doFunc()()
	return
}

func (_self *Delay) doFunc() func() {
	return func() {
		defer func() {
			if err := recover(); err != nil {
				_self.status.CompareAndSwap(Running, AfterRunningErr)
				return
			}
			if !_self.status.CompareAndSwap(Running, AfterRunningDone) {
				_self.status.Store(Cancel)
			}

		}()
		select {
		case <-time.After(_self.duration):
			if _self.status.CompareAndSwap(Wait, Running) {
				_self.do(_self.ctx)
			}
		case <-_self.ctx.Done():
			return
		}
	}

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
