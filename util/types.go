package util

import "reflect"

// IsImplements
// IsImplements(&struct{}, (*struct)(nil))
// check implements interface
/**

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
	println(IsImplements(Son1{}, (*Father)(nil))) >>false
	println(IsImplements(&Son2{}, (*Father)(nil))) >>true
	println(IsImplements(&FakeSon{}, (*Father)(nil))) >> false
}

*/
func IsImplements(obj interface{}, ifPtr interface{}) bool {
	objType := reflect.TypeOf(obj)
	tarElem := reflect.TypeOf(ifPtr).Elem()

	return objType.Implements(tarElem)
}
