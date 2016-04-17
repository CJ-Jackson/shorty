package http_error

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestErrorHandler(t *testing.T) {
	Convey("ErrorHandler", t, func() {
		var errorCode int

		_, errorCode = (ErrorHandler{R: nil}).getHttpHandler()
		So(errorCode, ShouldEqual, 0)

		_, errorCode = (ErrorHandler{
			R: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
		}).getHttpHandler()
		So(errorCode, ShouldEqual, 1)

		_, errorCode = (ErrorHandler{
			R: fmt.Errorf("Hi"),
		}).getHttpHandler()
		So(errorCode, ShouldEqual, 2)

		_, errorCode = (ErrorHandler{
			R: "Hi",
		}).getHttpHandler()
		So(errorCode, ShouldEqual, 3)
	})
}
