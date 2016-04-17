package http

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHttpBoot(t *testing.T) {
	Convey("callFnOnDebugOnly", t, func() {
		called := false

		fn := func() {
			called = true
		}

		(httpBoot{}).callFnOnDebugOnly(fn)

		So(called, ShouldBeFalse)

		(httpBoot{debug: true}).callFnOnDebugOnly(fn)

		So(called, ShouldBeTrue)
	})
}
