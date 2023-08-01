package util

import (
	"fmt"
	"reflect"
	"testing"
)

type Father interface {
	Father()
}
type Son1 struct {
	i int
}

func (_self *Son1) Father() {
	fmt.Println()
}

type Son2 struct {
	i int
}

func (_self *Son2) Father() {
	fmt.Println()

}

type FakeSon struct {
}

func TestIsImplementInterface(t *testing.T) {
	println(isImplements(&Son1{}, (*Father)(nil)))

	println(isImplements(&Son2{}, (*Father)(nil)))
	println(isImplements(&FakeSon{}, (*Father)(nil)))
}

func isImplements(obj interface{}, ifPtr interface{}) bool {
	objType := reflect.TypeOf(obj)
	tarElem := reflect.TypeOf(ifPtr).Elem()

	return objType.Implements(tarElem)
}
