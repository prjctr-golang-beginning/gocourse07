package main

import (
	"cycle_import/b"
	"cycle_import/sl"
)
import "cycle_import/a"

func main() {
	sl := sl.ServiceLocator{}
	sl.Register(b.CreateB())
	sl.Register(a.CreateA())

	b.InvokeSomethingFromA(sl)
	a.InvokeSomethingFromB(sl)
}
