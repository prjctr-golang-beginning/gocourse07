package main

import (
	"cycle_import/b"
	sl2 "cycle_import/sl"
)
import "cycle_import/a"

func main() {
	sl := sl2.ServiceLocator{}
	sl.Register(b.CreateB())
	sl.Register(a.CreateA())

	b.InvokeSomethingFromA(sl)
	a.InvokeSomethingFromB(sl)
}
