package skeleton

import (
	"bytes"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestSkeleton(t *testing.T) {
	Convey("Skeleton", t, func() {

		testSubject := &Skeleton{
			tpl: skeletonHtml,
		}

		testSubject.SetTitle("--title--")
		testSubject.SetHead(strings.NewReader("--head--"))
		testSubject.SetBody(strings.NewReader("--body--"))
		testSubject.SetFooter(strings.NewReader("--footer--"))
		testSubject.SetJavascript(strings.NewReader("--javascript--"))

		Convey("Check output of skeleton", func() {
			buf := &bytes.Buffer{}

			testSubject.execute(buf)

			output := buf.String()
			buf.Reset()

			So(strings.LastIndex(output, "--title--"), ShouldNotEqual, -1)
			So(strings.LastIndex(output, "--head--"), ShouldNotEqual, -1)
			So(strings.LastIndex(output, "--body--"), ShouldNotEqual, -1)
			So(strings.LastIndex(output, "--footer--"), ShouldNotEqual, -1)
			So(strings.LastIndex(output, "--javascript--"), ShouldNotEqual, -1)
		})
	})
}
