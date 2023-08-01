package cache

import (
	"context"
	"errors"
	"github.com/cafe-old-babe/glt/base"
	"strings"
	"sync"
	"time"
)

type SafeMemoryCache struct {
	dataMap       sync.Map
	errData       sync.Map
	syncMap       sync.Map
	IgnoreErrors  []error
	ErrorDuration time.Duration
}

// LoadOrStore  cache.GetValIgnoreNullKey not nil , allow gf null
func (_self *SafeMemoryCache) LoadOrStore(ctx context.Context, key any, gf base.CtxAFuncAE) (val any, err error) {

	var exists bool
	if val, exists = _self.dataMap.Load(key); exists {
		return
	}
	if val, exists = _self.errData.Load(key); exists {
		err = val.(error)
		return
	}

	if gf == nil {
		if nil == ctx.Value(GetValIgnoreNullKey) {
			err = errors.New("gf is nil")
		}
		return
	}

	actual, _ := _self.syncMap.LoadOrStore(key, new(sync.Mutex))
	mutex := actual.(*sync.Mutex)
	mutex.Lock()
	defer func() {
		mutex.Unlock()
		_self.syncMap.Delete(key)
	}()

	if val, exists = _self.dataMap.Load(key); exists {
		return
	}
	if val, exists = _self.errData.Load(key); exists {
		err = val.(error)
		return
	}
	defer func() {
		if rErr := recover(); rErr != nil {
			_self.storeErr(key, rErr.(error))
		}
	}()
	if val, err = gf(ctx, key); err != nil {
		_self.storeErr(key, err)
	} else {
		_self.Store(key, val)
	}
	return

}

func (_self *SafeMemoryCache) storeErr(key any, err error) {
	_self.errData.Store(key, err)

	if nil != _self.IgnoreErrors {
		for _, igErr := range _self.IgnoreErrors {
			if strings.Contains(err.Error(), igErr.Error()) {
				return
			}
		}
	}
	if _self.ErrorDuration <= 0 {
		return
	}
	go func(_t *SafeMemoryCache, k any) {
		select {
		case <-time.After(_t.ErrorDuration):
			_t.errData.Delete(k)
		}
	}(_self, key)

}

// AsyncClear
func (_self *SafeMemoryCache) AsyncClear(ctx context.Context, fs ...base.CtxAFunc) *sync.WaitGroup {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func(wg *sync.WaitGroup, _this *SafeMemoryCache, c context.Context) {
		defer wg.Done()
		_self.errData.Range(func(key, value any) bool {
			_self.errData.Delete(key)
			return true
		})
		_self.dataMap.Range(func(key, value any) (n bool) {
			defer func() {
				n = true
			}()
			_this.Delete(c, key, fs...)
			return
		})
		_self.syncMap.Range(func(key, value any) bool {
			_self.syncMap.Delete(key)
			return true
		})
	}(&waitGroup, _self, ctx)
	return &waitGroup
}

// Store
func (_self *SafeMemoryCache) Store(key any, val any) {
	_self.dataMap.Store(key, val)
}

// Range
func (_self *SafeMemoryCache) Range(f func(any, any) bool) {
	_self.dataMap.Range(func(key, value any) bool {
		return f(key, value)
	})
}

// Delete
func (_self *SafeMemoryCache) Delete(ctx context.Context, key any, fs ...base.CtxAFunc) {
	defer func() {
		_self.errData.Delete(key)
		_self.syncMap.Delete(key)
	}()
	value, loaded := _self.dataMap.LoadAndDelete(key)
	if !loaded {
		return
	}
	if nil != fs && len(fs) > 0 {
		for i := range fs {
			fs[i](ctx, value)
		}
	}
}

// Load
func (_self *SafeMemoryCache) Load(key any) (val any, ok bool) {
	return _self.dataMap.Load(key)
}
