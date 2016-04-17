package forms

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFunctions(t *testing.T) {
	Convey("ErrorToHtml", t, func() {
		So(ErrorToHtml(nil), ShouldBeEmpty)
		So(ErrorToHtml(fmt.Errorf("Hi")), ShouldEqual, `<span class="help-block">Hi</span>`)
	})
}
