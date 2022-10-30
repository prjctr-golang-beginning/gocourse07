package main

import (
	"fmt"
	"intrnl/pkg"
	"intrnl/pkg/internal"
)

func main() {
	p := pkg.SomePkg{`Some package`}
	fmt.Println(p)

	_ = internal.SomeEncapsulated{}
}
