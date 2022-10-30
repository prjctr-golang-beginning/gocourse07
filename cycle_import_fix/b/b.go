package b

import (
	"cycle_import/sl"
	"fmt"
)

type WithA interface {
	DoSomethingWithA()
}

type B struct {
	name string
}

func (b B) DoSomethingWithB() {
	fmt.Println(b)
}
func CreateB() *B {
	b := B{`I am B`}
	return &b
}
func InvokeSomethingFromA(sl sl.ServiceLocator) {
	var o WithA
	sl.Get(&o)
	o.DoSomethingWithA()
}
