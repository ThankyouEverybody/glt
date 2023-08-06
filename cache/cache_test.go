package cache

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestCache(t *testing.T) {
	es := make([]error, 0)
	es = append(es, errors.New("abc"))
	newCache := &SafeMemoryCache[string, string]{
		IgnoreErrors:  es,
		ErrorDuration: 0,
	}
	var wait sync.WaitGroup
	for i := 0; i < 100; i++ {
		for t := 0; t < 10; t++ {
			wait.Add(1)
			go func(j int) {
				Key := strconv.Itoa(j)
				_, _ = newCache.LoadOrStore(context.TODO(), Key,
					func(ctx context.Context, key string) (v string, err error) {
						fmt.Printf("gf ---> %v\n", key)
						return key + " val", nil
					})
				wait.Done()
			}(t)
		}
		for t := 9; t >= 0; t-- {
			wait.Add(1)
			go func(j int) {
				Key := strconv.Itoa(j)
				_, _ = newCache.LoadOrStore(context.TODO(), Key,
					func(ctx context.Context, key string) (v string, err error) {
						fmt.Printf("gf ---> %v\n", key)
						return key + " val", nil
					})
				wait.Done()
			}(t)
		}
	}
	wait.Wait()
	newCache.Store("abc", "abcVal")
	newCache.Range(func(k, v string) bool {
		fmt.Printf("k: %s, v: %s\n", k, v)
		return true
	})
	newCache.Delete(context.TODO(), "abc", func(ctx context.Context, val string) {
		fmt.Printf("delete abc val: %s\n", val)
	})
}
