package a

import (
	"cycle_import/b"
	"fmt"
)

type WithB interface {
	DoSomethingWithB()
}

type A struct {
	name string
}

func (a A) DoSomethingWithA() {
	fmt.Println(a)
}
func CreateA() *A {
	a := A{`I am A`}
	return &a
}
func InvokeSomethingFromB( /*sl sl.ServiceLocator*/ ) {
	o := b.CreateB()
	//var o WithB
	//sl.Get(&o)
	o.DoSomethingWithB()
}
