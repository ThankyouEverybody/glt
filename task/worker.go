package task

import (
	"context"
	"errors"
	"github.com/ThankyouEverybody/glt/base"
	"sync"
	"sync/atomic"
)

type Worker struct {
	key     string
	doQ     chan base.FuncCtxE
	errFunc base.CtxEFunc
	wait    sync.WaitGroup
	state   *atomic.Int32
}

func NewWorker(key string, ef base.CtxEFunc) (*Worker, error) {
	return NewNumWorker(key, ef, 1024)
}

func NewNumWorker(key string, ef base.CtxEFunc, num int) (*Worker, error) {
	if len(key) <= 0 || num <= 0 {
		return nil, errors.New("arg is invalid")
	}
	var state atomic.Int32
	state.Store(0)

	w := &Worker{
		key:     key,
		doQ:     make(chan base.FuncCtxE, num),
		errFunc: ef,
		state:   &state,
	}
	go w.start()
	return w, nil
}

func (_self *Worker) Do(f base.FuncCtxE) {
	_self.doQ <- func() (ctx context.Context, err error) {
		defer func() {
			if re := recover(); re != nil {
				err = re.(error)
				ctx = context.TODO()
				return
			}
		}()
		return f()
	}
}

func (_self *Worker) start() {
	go func(w *Worker) {

		for {
			select {
			case f, ok := <-w.doQ:
				if !ok {
					w.wait.Done()
					return
				}
				if ctx, err := f(); err != nil && nil != w.errFunc {
					w.errFunc(ctx, err)
				}
			}
		}
	}(_self)

}

func (_self *Worker) DoOver() (err error) {
	if _self.state.CompareAndSwap(0, 1) {
		_self.wait.Add(1)
		close(_self.doQ)
	} else {
		err = errors.New("DoOver duplicate executed")
	}
	return

}

func (_self *Worker) DoWait() {
	_ = _self.DoOver()
	_self.wait.Wait()
	return
}
