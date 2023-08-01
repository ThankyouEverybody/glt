package util

import (
	"fmt"
	"testing"
)

func TestCovers(t *testing.T) {
	type testStruct struct {
		a int
		b float64
	}
	test := testStruct{
		a: 1,
		b: 1.1,
	}
	var a int = 10
	var b int8 = 10
	var c int16 = 10
	var d int32 = 10
	var e int64 = 10
	var f uint8 = 10
	var g uint16 = 10
	var h uint32 = 10
	var i uint64 = 10
	var j float32 = 10.90
	var k float64 = 10.90
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", c)
	fmt.Printf("%d\n", d)
	fmt.Printf("%d\n", e)
	fmt.Printf("%d\n", f)
	fmt.Printf("%d\n", g)
	fmt.Printf("%d\n", h)
	fmt.Printf("%d\n", i)
	fmt.Printf("%f\n", j)
	fmt.Printf("%f\n", k)
	fmt.Printf("----------\n")

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", d)
	fmt.Printf("%#v\n", e)
	fmt.Printf("%#v\n", f)
	fmt.Printf("%#v\n", g)
	fmt.Printf("%#v\n", h)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%#v\n", j)
	fmt.Printf("%#v\n", k)
	fmt.Printf("%#v\n", test)
	fmt.Printf("----------\n")

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", g)
	fmt.Printf("%v\n", h)
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", j)
	fmt.Printf("%v\n", k)
	fmt.Printf("%v\n", test)

	fmt.Printf("----------\n")

	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
	fmt.Printf("%+v\n", c)
	fmt.Printf("%+v\n", d)
	fmt.Printf("%+v\n", e)
	fmt.Printf("%+v\n", f)
	fmt.Printf("%+v\n", g)
	fmt.Printf("%+v\n", h)
	fmt.Printf("%+v\n", i)
	fmt.Printf("%+v\n", j)
	fmt.Printf("%+v\n", k)
	fmt.Printf("%+v\n", test)

}

func TestCovers2(t *testing.T) {
	type testStruct struct {
		a int
		b float64
	}
	test := testStruct{
		a: 1,
		b: 1.1,
	}
	var a = []int{10}
	var b = []int8{10}
	var c = []int16{10}
	var d = []int32{10}
	var e = []int64{10}
	var f = []uint8{10}
	var g = []uint16{10}
	var h = []uint32{10}
	var i = []uint64{10}
	var j = []float32{10.90}
	var k = []float64{10.90}
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", c)
	fmt.Printf("%d\n", d)
	fmt.Printf("%d\n", e)
	fmt.Printf("%d\n", f)
	fmt.Printf("%d\n", g)
	fmt.Printf("%d\n", h)
	fmt.Printf("%d\n", i)
	fmt.Printf("%f\n", j)
	fmt.Printf("%f\n", k)
	fmt.Printf("----------\n")

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", d)
	fmt.Printf("%#v\n", e)
	fmt.Printf("%#v\n", f)
	fmt.Printf("%#v\n", g)
	fmt.Printf("%#v\n", h)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%#v\n", j)
	fmt.Printf("%#v\n", k)
	fmt.Printf("%#v\n", test)
	fmt.Printf("----------\n")

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", g)
	fmt.Printf("%v\n", h)
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", j)
	fmt.Printf("%v\n", k)
	fmt.Printf("%v\n", test)

	fmt.Printf("----------\n")

	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
	fmt.Printf("%+v\n", c)
	fmt.Printf("%+v\n", d)
	fmt.Printf("%+v\n", e)
	fmt.Printf("%+v\n", f)
	fmt.Printf("%+v\n", g)
	fmt.Printf("%+v\n", h)
	fmt.Printf("%+v\n", i)
	fmt.Printf("%+v\n", j)
	fmt.Printf("%+v\n", k)
	fmt.Printf("%+v\n", test)

}
