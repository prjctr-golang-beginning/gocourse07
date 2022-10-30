package main

import (
	"cycle_import/b"
)
import "cycle_import/a"

func main() {
	b.InvokeSomethingFromA()
	a.InvokeSomethingFromB()
}
