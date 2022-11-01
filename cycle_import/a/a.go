package a

import (
	"fmt"
)

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
func InvokeSomethingFromB(sl Getter) {
	var b Ber
	sl.Get(&b)
	b.DoSomethingWithB()
}

type Ber interface {
	DoSomethingWithB()
}

type Getter interface {
	Get(some any) bool
}
