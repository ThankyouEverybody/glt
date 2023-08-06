package task

import (
	"context"
	"errors"
	"github.com/cafe-old-babe/glt/base"
	"github.com/cafe-old-babe/glt/container/queue"
	"sync"
	"time"
)

type DelayPool struct {
	delayTaskQueue *queue.PriorityQueue[Delay]
	worker         *Worker
	ctx            context.Context
	destroyNotify  context.CancelFunc
	popNotify      context.CancelFunc
	putNotify      context.CancelFunc
	lock           sync.Mutex
}

func NewDelayPool() (*DelayPool, error) {
	priorityQueue, err := queue.NewPriorityQueue[Delay](func(a, b Delay) int {
		return int(a.duration - b.duration)
	})
	if err != nil {
		return nil, err
	}
	worker, err := NewWorker("DelayPool", nil)
	if err != nil {
		return nil, err
	}

	ctx, cancelFunc := context.WithCancel(context.TODO())
	dp := &DelayPool{
		delayTaskQueue: priorityQueue,
		worker:         worker,
		ctx:            ctx,
		destroyNotify:  cancelFunc,
	}
	go dp.pop()
	return dp, nil
}

func (_self *DelayPool) Put(duration time.Duration, f base.CtxFunc,
	ctx ...context.Context) (dt *Delay, err error) {

	if duration <= 0 {
		err = errors.New("duration then less and equals zero")
		return
	}
	if nil == f {
		err = errors.New("do func is nil")
		return
	}

	dt, err = newDelayPtr(duration, f, ctx...)
	if err != nil {
		return
	}
	_self.lock.Lock()
	defer _self.lock.Unlock()
	_self.delayTaskQueue.Push(dt)

	if _self.popNotify != nil {
		_self.popNotify()
	} else if _self.putNotify != nil {
		_self.putNotify()
	}
	return

}

func (_self *DelayPool) Cancel(delay *Delay) error {
	if delay == nil {
		return errors.New("arg is nil")
	}
	_self.lock.Lock()
	defer _self.lock.Unlock()
	ok := _self.delayTaskQueue.Delete(delay)
	if !ok {
		return errors.New("cancel failure, delay task not found")

	}
	if err := delay.Cancel(); err != nil {
		return err
	}
	return nil
}

func (_self *DelayPool) pop() {
	var delay *Delay
	var ok bool
	for {
		if !_self.lock.TryLock() {
			continue
		}
		//pop
		if delay, ok = _self.delayTaskQueue.Pop(); !ok {
			_self.lock.Unlock()
			popCtx, cancel := context.WithCancel(_self.ctx)
			_self.popNotify = cancel
			select {
			case <-popCtx.Done():
				_self.popNotify = nil
			case <-_self.ctx.Done():
				_self.destroy(delay)
				return
			}
		}
		_self.lock.Unlock()
		nextCtx, nextCancel := context.WithCancel(_self.ctx)
		_self.worker.Do(func() (context.Context, error) {
			defer func() {
				if delay.Status() != Wait {
					nextCancel()
				}
			}()
			delay.doFunc()()
			return nil, nil
		})

		putCtx, putCancel := context.WithCancel(_self.ctx)
		_self.putNotify = putCancel
		select {
		case <-nextCtx.Done():
			continue
		case <-putCtx.Done():
			delay.cancel()
			if delay.Status() == Wait {
				_self.delayTaskQueue.Push(delay)
			}
		case <-_self.ctx.Done():
			_self.destroy(delay)
			return
		}

	}

}

func (_self *DelayPool) destroy(delay *Delay) {
	if delay != nil {
		delay.cancel()
		delay = nil
	}
	_self.delayTaskQueue.Destroy()
	_self.worker.DoWait()
}
