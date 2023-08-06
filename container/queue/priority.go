package queue

import (
	_ "container/heap"
	"errors"
	"sync"
)

type PriorityQueue[T any] struct {
	heap []*T
	// a < b return -1; a > b return 1; a == b return 0
	// -1、0、1，表示小于、等于、大于
	compare func(T, T) int
	lock    sync.Mutex
}

// NewPriorityQueue 创建一个新的优先级队列
// compare: 用于定义元素之间的优先级比较规则的函数，返回值为 -1、0、1，表示小于、等于、大于
// 返回值:
// - *PriorityQueue: 新创建的优先级队列实例指针
// - error: 如果 compare 为 nil，则返回错误
func NewPriorityQueue[T any](compare func(T, T) int) (*PriorityQueue[T], error) {
	if compare == nil {
		return nil, errors.New("compare is nil")
	}
	return &PriorityQueue[T]{
		heap:    make([]*T, 0, 1024),
		compare: compare,
	}, nil
}

// Push 向队列中插入元素
// data: 元素的数据信息
func (_self *PriorityQueue[T]) Push(data *T) {
	_self.lock.Lock()
	defer _self.lock.Unlock()
	_self.heap = append(_self.heap, data)
	_self.siftUp(len(_self.heap) - 1)
}

// Pop 删除并返回队列中的最高优先级元素
// 返回值:
// - error: 如果队列为空，则返回错误
func (_self *PriorityQueue[T]) Pop() (data *T, ok bool) {
	_self.lock.Lock()
	defer _self.lock.Unlock()
	if _self.IsEmpty() {
		return
	}
	ok = true
	data = _self.heap[0]
	if len(_self.heap) > 1 {
		last := _self.heap[len(_self.heap)-1]
		_self.heap = _self.heap[:len(_self.heap)-1]
		_self.heap[0] = last
		_self.siftDown(0)
	} else {
		_self.heap = _self.heap[:len(_self.heap)-1]
	}

	return
}

// IsEmpty 检查队列是否为空
// 返回值:
// - bool: 如果队列为空，则返回 true；否则返回 false
func (_self *PriorityQueue[T]) IsEmpty() bool {
	return _self.Size() == 0
}

// siftUp 上浮操作，用于维护堆的性质
// 比较当前元素与其父节点的优先级来确定是否需要交换位置，直到满足堆的性质为止
// index: 需要上浮的元素的索引
func (_self *PriorityQueue[T]) siftUp(index int) {
	parent := (index - 1) / 2
	for index > 0 && _self.compare(*_self.heap[index], *_self.heap[parent]) < 0 {
		_self.heap[index], _self.heap[parent] = _self.heap[parent], _self.heap[index]
		index = parent
		parent = (index - 1) / 2
	}
}

// siftDown 下沉操作，用于维护堆的性质
// 比较当前元素与其左右子节点的优先级来确定是否需要交换位置，直到满足堆的性质为止。
// index: 需要下沉的元素的索引
func (_self *PriorityQueue[T]) siftDown(index int) {
	size := len(_self.heap)
	leftChild := 2*index + 1
	rightChild := 2*index + 2
	smallest := index

	if leftChild < size && _self.compare(*_self.heap[leftChild], *_self.heap[smallest]) < 0 {
		smallest = leftChild
	}

	if rightChild < size && _self.compare(*_self.heap[rightChild], *_self.heap[smallest]) < 0 {
		smallest = rightChild
	}

	if smallest != index {
		_self.heap[index], _self.heap[smallest] = _self.heap[smallest], _self.heap[index]
		_self.siftDown(smallest)
	}
}

func (_self *PriorityQueue[T]) find(data *T) int {
	return _self.binarySearch(data, 0, len(_self.heap)-1)
}

func (_self *PriorityQueue[T]) binarySearch(data *T, start, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	compareResult := _self.compare(*_self.heap[mid], *data)
	if compareResult == 0 {
		if _self.heap[mid] == data {

			return mid
		}
	}
	if compareResult < 0 {
		return _self.binarySearch(data, mid+1, end)
	} else {
		return _self.binarySearch(data, start, mid-1)
	}
}

func (_self *PriorityQueue[T]) Delete(data *T) bool {

	_self.lock.Lock()
	defer _self.lock.Unlock()
	index := _self.find(data)
	if index < 0 {
		return false
	}
	size := len(_self.heap)
	if index < 0 || index >= size {
		return false
	}
	_self.heap[index] = _self.heap[size-1]
	_self.heap = _self.heap[:size-1]
	_self.siftDown(index)
	return true
}

func (_self *PriorityQueue[T]) Size() int {
	return len(_self.heap)
}

func (_self *PriorityQueue[T]) Destroy() {
	for i := 0; i < len(_self.heap); i++ {
		_self.heap[i] = nil
	}
	_self.heap = nil
	_self.compare = nil
}
