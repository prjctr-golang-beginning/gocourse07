package main

import (
	"cycle_import/a"
	"cycle_import/b"
	"cycle_import/sl"
)

func main() {
	sl := &sl.ServiceLocator{}
	sl.Register(a.CreateA())
	sl.Register(b.CreateB())
	b.InvokeSomethingFromA(sl)
	a.InvokeSomethingFromB(sl)
}
