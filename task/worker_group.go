package task

import (
	"context"
	"errors"
	"fmt"
	"github.com/ThankyouEverybody/glt/base"
	"sync"
	"sync/atomic"
)

type WorkerGroup struct {
	ctx            context.Context
	wait           sync.WaitGroup
	workerMap      sync.Map
	workerChanSize int
	hashWorker     base.CtxFuncSE
	errFunc        base.CtxEFunc
	state          *atomic.Int32
}

var defaultHashWorker base.CtxFuncSE = func(ctx context.Context) (string, error) {
	fmt.Printf("use default hashWorker\n")
	return "default", nil
}
var defaultErrFunc base.CtxEFunc = func(ctx context.Context, err error) {
	fmt.Printf("use default errFunc, err:%+v\n", err)
}

func NewWorkerGroup(ctx context.Context, workerChanSize int,
	hashWorker base.CtxFuncSE, errFunc base.CtxEFunc) (workerGroup *WorkerGroup, err error) {
	if workerChanSize <= 0 {
		err = errors.New("arg is invalid")
		return
	}
	workerGroup = new(WorkerGroup)
	workerGroup.ctx = ctx
	workerGroup.workerChanSize = workerChanSize
	if hashWorker == nil {
		workerGroup.hashWorker = defaultHashWorker
	} else {
		workerGroup.hashWorker = hashWorker
	}
	if errFunc == nil {
		workerGroup.errFunc = defaultErrFunc
	} else {
		workerGroup.errFunc = errFunc
	}
	var state atomic.Int32
	state.Store(0)
	workerGroup.state = &state
	go func(c context.Context) {
		select {
		case <-c.Done():
			fmt.Printf("done-->%+v", c.Err())
			workerGroup.Wait()
			fmt.Printf("wait done")
			return
		}
	}(ctx)
	return
}

func (_self *WorkerGroup) Push(ctx context.Context, f base.FuncCtxE) (err error) {
	if _self.state.Load() == 1 {
		err = errors.New("workerGroup is close")
		return
	}

	var workerKey string
	workerKey, err = _self.hashWorker(ctx)
	if err != nil {
		return
	}
	actual, loaded := _self.workerMap.LoadOrStore(workerKey, make(chan base.FuncCtxE, _self.workerChanSize))
	ch := actual.(chan base.FuncCtxE)
	if !loaded {
		_self.wait.Add(1)
		go func(key string, _this *WorkerGroup) {
			worker, _ := NewNumWorker(key, _this.errFunc, _this.workerChanSize)
			for {
				select {
				case f, ok := <-ch:
					if !ok {
						worker.DoWait()
						_this.workerMap.Delete(key)
						_this.wait.Done()
						return
					}
					worker.Do(f)
				}
			}
		}(workerKey, _self)
	}
	ch <- f
	return
}

func (_self *WorkerGroup) PushOver() (err error) {
	if _self.state.CompareAndSwap(0, 1) {
		_self.workerMap.Range(func(key, value any) bool {
			ch := value.(chan base.FuncCtxE)
			close(ch)
			return true
		})
	} else {
		err = errors.New("PushOver duplicate executed")
	}
	return
}

func (_self *WorkerGroup) Wait() {
	_ = _self.PushOver()
	_self.wait.Wait()
}
