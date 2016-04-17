package arguments

import (
	"fmt"
	"time"
)

/*
Implement:
	ArgumentTransformerInterface in "github.com/cjtoolkit/cli"
*/
type Time struct {
	Ptr     *time.Time // Mandatory
	Format  string     // Mandatory
	Min     time.Time
	MinZero bool
	Max     time.Time
	MaxZero bool
}

func (t Time) PreCheck() {
	switch {
	case nil == t.Ptr:
		panic("Ptr cannot be nil")
	case "" == t.Format:
		panic("Format cannot be empty")
	}
}

func (t Time) Constraint() string {
	str := fmt.Sprintf("Type:'time.Time' Format:'%s' Default:'%s'", t.Format, (*t.Ptr).Format(t.Format))

	if t.MinZero || !t.Min.IsZero() {
		str += fmt.Sprint(" Min:'", t.Min.Format(t.Format), "'")
	}

	if t.MaxZero || !t.Max.IsZero() {
		str += fmt.Sprint(" Max:'", t.Max.Format(t.Format), "'")
	}

	return str
}

func (t Time) ArgumentTransform(argument string) {
	t.populatePointer(argument)
	t.validate()
}

func (t Time) populatePointer(value string) {
	parsedTime, err := time.ParseInLocation(t.Format, value, time.Local)
	if nil != err {
		panic("Unable to parse time, probably didn't match the format")
	}

	*t.Ptr = parsedTime
}

func (t Time) validate() {
	t.validateMin()
	t.validateMax()
}

func (t Time) validateMin() {
	switch {
	case t.Min.IsZero() && !t.MinZero:
		return
	case (*t.Ptr).Unix() < t.Min.Unix():
		panic(fmt.Sprintf("Should be more than '%s'", t.Min.Format(t.Format)))
	}
}

func (t Time) validateMax() {
	switch {
	case t.Max.IsZero() && !t.MaxZero:
		return
	case (*t.Ptr).Unix() > t.Max.Unix():
		panic(fmt.Sprintf("Should be less than '%s'", t.Max.Format(t.Format)))
	}
}
