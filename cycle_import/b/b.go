package b

import (
	"fmt"
)

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
func InvokeSomethingFromA(sl Getter) {
	var a Aer
	sl.Get(&a)
	a.DoSomethingWithA()
}

type Aer interface {
	DoSomethingWithA()
}

type Getter interface {
	// ;lfjgnb'aoeinr[ q034ijg]09q34hg0 [q3i4
	Get(some any) bool
}

type MyGetter struct {
}
