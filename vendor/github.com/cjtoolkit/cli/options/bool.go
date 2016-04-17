package options

import (
	"fmt"
	"github.com/cjtoolkit/cli"
)

/*
Implement:
	OptionTransformerInterface in "github.com/cjtoolkit/cli"
*/
type Bool struct {
	Ptr *bool // Mandatory
}

func (b Bool) PreCheck() {
	if nil == b.Ptr {
		panic("Ptr cannot be nil")
	}
}

func (b Bool) Constraint() string {
	return "Type:'bool' Default:'" + fmt.Sprint(*b.Ptr) + "'"
}

func (b Bool) OptionTransform(option cli.OptionsInterface) {
	*b.Ptr = !*b.Ptr
}
