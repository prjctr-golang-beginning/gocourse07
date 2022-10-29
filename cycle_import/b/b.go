package b

import (
	"cycle_import/a"
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
func InvokeSomethingFromA( /*sl sl.ServiceLocator*/ ) {
	o := a.CreateA()
	//var o WithA
	//sl.Get(&o)
	o.DoSomethingWithA()
}
