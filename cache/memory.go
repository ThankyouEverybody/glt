package cache

import (
	"context"
	"errors"
	"github.com/cafe-old-babe/glt/base"
	"strings"
	"sync"
	"time"
)

type SafeMemoryCache[K any, V any] struct {
	dataMap       sync.Map
	errData       sync.Map
	syncMap       sync.Map
	IgnoreErrors  []error
	ErrorDuration time.Duration
}

// LoadOrStore  cache.GetValIgnoreNullKey not nil , allow gf null
func (_self *SafeMemoryCache[K, V]) LoadOrStore(ctx context.Context, key K, gf base.CtxPFuncRE[K, V]) (val V, err error) {

	var exists bool
	var loadVal any
	var errVal any
	if loadVal, exists = _self.dataMap.Load(key); exists {
		val = loadVal.(V)
		return
	}
	if errVal, exists := _self.errData.Load(key); exists {
		err = errVal.(error)
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

	if loadVal, exists = _self.dataMap.Load(key); exists {
		val = loadVal.(V)
		return
	}
	if errVal, exists = _self.errData.Load(key); exists {
		err = errVal.(error)
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

func (_self *SafeMemoryCache[K, V]) storeErr(key any, err error) {
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
	go func(_t *SafeMemoryCache[K, V], k any) {
		select {
		case <-time.After(_t.ErrorDuration):
			_t.errData.Delete(k)
		}
	}(_self, key)

}

// AsyncClear
func (_self *SafeMemoryCache[K, V]) AsyncClear(ctx context.Context, fs ...base.CtxTFunc[V]) *sync.WaitGroup {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func(wg *sync.WaitGroup, _this *SafeMemoryCache[K, V], c context.Context) {
		defer wg.Done()
		_self.errData.Range(func(key, value any) bool {
			_self.errData.Delete(key.(K))
			return true
		})
		_self.dataMap.Range(func(key, value any) (n bool) {
			defer func() {
				n = true
			}()
			_this.Delete(c, key.(K), fs...)
			return
		})
		_self.syncMap.Range(func(key, value any) bool {
			_self.syncMap.Delete(key.(K))
			return true
		})
	}(&waitGroup, _self, ctx)
	return &waitGroup
}

// Store
func (_self *SafeMemoryCache[K, V]) Store(key K, val V) {
	_self.dataMap.Store(key, val)
}

// Range
func (_self *SafeMemoryCache[K, V]) Range(f func(K, V) bool) {
	_self.dataMap.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

// Delete
func (_self *SafeMemoryCache[K, V]) Delete(ctx context.Context, key K, fs ...base.CtxTFunc[V]) {
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
			fs[i](ctx, value.(V))
		}
	}
}

// Load
func (_self *SafeMemoryCache[K, V]) Load(key K) (val V, ok bool) {
	var loadVal any
	if loadVal, ok = _self.dataMap.Load(key); ok {
		val = loadVal.(V)
	}

	return
}
