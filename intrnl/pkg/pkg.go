package pkg

import (
	"fmt"
	"intrnl/pkg/internal"
)

type SomePkg struct {
	Name string
}

func (sp SomePkg) String() string {
	_ = internal.SomeEncapsulated{}
	return fmt.Sprintf(`My name is %s`, sp.Name)
}
