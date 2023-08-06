package set

import (
	"testing"
)

type TestSetStruct struct {
	a int
}

func TestSet(t *testing.T) {
	set := NewSet[TestSetStruct]()
	t.Logf("set is empty: %v\n", set.IsEmpty())
	t.Logf("set len: %d\n", set.Size())
	var last *TestSetStruct
	for i := 0; i < 10; i++ {
		last = &TestSetStruct{a: i}
		set.Add(last)
	}

	t.Logf("set is empty: %v\n", set.IsEmpty())
	t.Logf("set len: %d\n", set.Size())

	t.Logf("set contains last %v\n", set.Contains(last))
	set.Remove(last)
	t.Logf("set is empty: %v\n", set.IsEmpty())
	t.Logf("set len: %d\n", set.Size())
	t.Logf("set contains last %v\n", set.Contains(last))

	last = &TestSetStruct{a: 0}
	t.Logf("set contains last %v\n", set.Contains(last))
	set.Clear()
	t.Logf("set is empty: %v\n", set.IsEmpty())
	t.Logf("set len: %d\n", set.Size())

	newSet := NewSet[TestSetStruct]()
	newSet.Add(last)
	for i := 10; i < 20; i++ {
		newSet.Add(&TestSetStruct{a: i})
	}
	newSet.Add(last)

	t.Logf("newSet is empty: %v\n", newSet.IsEmpty())
	t.Logf("newSet len: %d\n", newSet.Size())

	set.Merge(newSet)

	t.Logf("set is empty: %v\n", set.IsEmpty())
	t.Logf("set len: %d\n", set.Size())
	t.Logf("set range\n")

	set.Range(func(ptr *TestSetStruct) bool {
		t.Logf("%v\n", ptr.a)
		return true
	})
	t.Logf("set range\n")

}
