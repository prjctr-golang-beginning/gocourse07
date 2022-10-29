package main

import (
	"cycle_import/b"
)
import "cycle_import/a"

func main() {
	//sl := sl2.ServiceLocator{}
	//sl.Register(b.CreateB())
	//sl.Register(a.CreateA())

	b.InvokeSomethingFromA( /*sl*/ )
	a.InvokeSomethingFromB( /*sl*/ )
}
