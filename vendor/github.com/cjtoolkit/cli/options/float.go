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
type Float struct {
	Ptr     *float64 // Mandatory
	Min     float64
	MinZero bool
	Max     float64
	MaxZero bool
}

func (f Float) PreCheck() {
	if nil == f.Ptr {
		panic("Ptr cannot be nil")
	}
}

func (f Float) Constraint() string {
	str := "Type:'float64' Default:'" + fmt.Sprint(*f.Ptr) + "'"

	if f.MinZero || 0 != f.Min {
		str += fmt.Sprint(" Min:'", f.Min, "'")
	}

	if f.MaxZero || 0 != f.Max {
		str += fmt.Sprint(" Max:'", f.Max, "'")
	}

	return str
}

func (f Float) OptionTransform(option cli.OptionsInterface) {
	f.populatePointer(option.GetOne())
	f.validate()
}

func (f Float) populatePointer(value string) {
	num, err := strconv.ParseFloat(value, FLOAT_BIT)
	if nil != err {
		panic("Not a float")
	}

	*f.Ptr = num
}

func (f Float) validate() {
	f.validateMin()
	f.validateMax()
}

func (f Float) validateMin() {
	switch {
	case 0 == f.Min && !f.MinZero:
		return
	case *f.Ptr < f.Min:
		panic(fmt.Sprintf("Should be more than '%g'", f.Min))
	}
}

func (f Float) validateMax() {
	switch {
	case 0 == f.Max && !f.MaxZero:
		return
	case *f.Ptr > f.Max:
		panic(fmt.Sprintf("Should be less than '%g'", f.Max))
	}
}
