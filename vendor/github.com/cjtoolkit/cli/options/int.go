package options

import (
	"fmt"
	"github.com/cjtoolkit/cli"
	"strconv"
)

/*
Implement:
	OptionTransformerInterface in "github.com/cjtoolkit/cli"
*/
type Int struct {
	Ptr     *int64 // Mandatory
	Min     int64
	MinZero bool
	Max     int64
	MaxZero bool
}

func (i Int) PreCheck() {
	if nil == i.Ptr {
		panic("Ptr cannot be nil")
	}
}

func (i Int) Constraint() string {
	str := "Type:'int64' Default:'" + fmt.Sprint(*i.Ptr) + "'"

	if i.MinZero || 0 != i.Min {
		str += fmt.Sprint(" Min:'", i.Min, "'")
	}

	if i.MaxZero || 0 != i.Max {
		str += fmt.Sprint(" Max:'", i.Max, "'")
	}

	return str
}

func (i Int) OptionTransform(option cli.OptionsInterface) {
	i.populatePointer(option.GetOne())
	i.validate()
}

func (i Int) populatePointer(value string) {
	num, err := strconv.ParseInt(value, INT_DINARY, INT_BIT)
	if nil != err {
		panic("Not an integer")
	}

	*i.Ptr = num
}

func (i Int) validate() {
	i.validateMin()
	i.validateMax()
}

func (i Int) validateMin() {
	switch {
	case 0 == i.Min && !i.MinZero:
		return
	case *i.Ptr < i.Min:
		panic(fmt.Sprintf("Should be more than '%d'", i.Min))
	}
}

func (i Int) validateMax() {
	switch {
	case 0 == i.Max && !i.MaxZero:
		return
	case *i.Ptr > i.Max:
		panic(fmt.Sprintf("Should be less than '%d'", i.Max))
	}
}
