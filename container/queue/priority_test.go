package queue

import (
	"fmt"
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	pq, err := NewPriorityQueue[float64](func(a, b float64) int {
		f := a - b
		if f < 0 {
			return -1
		} else if f == 0 {
			return 0
		}
		return 1
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for !pq.IsEmpty() {
		elem, _ := pq.Pop()
		fmt.Printf("Data: %v\n", elem)
	}

}
