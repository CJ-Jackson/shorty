package common

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

type readerWithAnxiety struct{}

func (_ readerWithAnxiety) Read(p []byte) (n int, err error) {
	err = fmt.Errorf("I've anxiety")
	return
}

func TestFunctions(t *testing.T) {
	chPanic := make(chan interface{})

	panicTrap := func(fn func()) {
		defer func() {
			chPanic <- recover()
		}()
		fn()
	}

	Convey("CheckError", t, func() {
		go panicTrap(func() {
			CheckError(nil)
		})

		So(<-chPanic, ShouldBeNil)

		err := fmt.Errorf("Hi")

		go panicTrap(func() {
			CheckError(err)
		})

		So(<-chPanic, ShouldEqual, err)
	})

	Convey("PanicIfNotNil", t, func() {
		go panicTrap(func() {
			PanicIfNotNil(nil)
		})

		So(<-chPanic, ShouldBeNil)

		err := fmt.Errorf("Hi")

		go panicTrap(func() {
			PanicIfNotNil(err)
		})

		So(<-chPanic, ShouldEqual, err)
	})

	Convey("ReaderToHtml", t, func() {
		var i int

		_, i = renderToHtml(nil)
		So(i, ShouldEqual, 1)

		_, i = renderToHtml(readerWithAnxiety{})
		So(i, ShouldEqual, 2)

		_, i = renderToHtml(strings.NewReader("Hi"))
		So(i, ShouldEqual, 0)
	})

	Convey("TrueFalseExecFunction", t, func() {
		str := ""

		trueFn := func() {
			str = "a"
		}

		falseFn := func() {
			str = "b"
		}

		TrueFalseExecFunction(false, trueFn, falseFn)
		So(str, ShouldEqual, "b")

		TrueFalseExecFunction(true, trueFn, falseFn)
		So(str, ShouldEqual, "a")
	})
}
