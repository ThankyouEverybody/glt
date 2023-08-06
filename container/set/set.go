package set

import "errors"

type Set[T any] struct {
	m map[*T]any
}

func GenerateNewSet[T any](slice []interface{}, f func(interface{}) *T) (*Set[T], error) {
	if f == nil {
		return nil, errors.New("f is nil")
	}
	newSet := NewSet[T]()
	if nil != slice && len(slice) > 0 {
		for _, data := range slice {
			newSet.Add(f(data))
		}
	}
	return newSet, nil
}

func NewSet[T any]() *Set[T] {
	return &Set[T]{
		m: make(map[*T]any),
	}
}

func (_self *Set[T]) Add(t *T) bool {
	if t == nil {
		return false
	}
	if !_self.Contains(t) {
		_self.m[t] = nil
	}
	return true
}

func (_self *Set[T]) Contains(t *T) bool {
	_, ok := _self.m[t]
	return ok
}

func (_self *Set[T]) Remove(t *T) {
	delete(_self.m, t)
}

func (_self *Set[T]) Size() int {
	return len(_self.m)
}

func (_self *Set[T]) IsEmpty() bool {
	return _self.Size() == 0
}

func (_self *Set[T]) Range(f func(*T) bool) {

	if nil == _self || _self.IsEmpty() {
		return
	}
	for data, _ := range _self.m {
		if !f(data) {
			return
		}
	}
}

func (_self *Set[T]) Merge(other *Set[T]) {
	if other == nil || other.IsEmpty() {
		return
	}
	other.Range(func(t *T) bool {
		_self.Add(t)
		return true
	})
}

func (_self *Set[T]) Clear() {

	_self.Range(func(t *T) bool {
		delete(_self.m, t)
		return true
	})
}
