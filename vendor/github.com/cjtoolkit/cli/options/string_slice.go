package options

import (
	"fmt"
	"github.com/cjtoolkit/cli"
)

type StringSlice struct {
	Ptr *[]string
}

func (sS StringSlice) PreCheck() {
	if nil == sS.Ptr {
		panic("Ptr cannot be nil")
	}
}

func (sS StringSlice) Constraint() string {
	return "Type:'[]string' Default:'" + fmt.Sprint(*sS.Ptr) + "'"
}

func (sS StringSlice) OptionTransform(option cli.OptionsInterface) {
	*sS.Ptr = option.GetAll()
}
