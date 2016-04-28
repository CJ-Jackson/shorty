package forms

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestUrlForm(t *testing.T) {
	Convey("UrlForm", t, func() {

		testSubject := &UrlForm{
			tpl:      urlFormHtml,
			UrlModel: "http://www.example.com",
			UrlErr:   fmt.Errorf("I am UrlError"),
			UrlNorm:  "http://www.example.com",
		}

		testSubject.Fields()

		Convey("Test Output", func() {
			output := string(testSubject.GetHtml())

			So(strings.Index(output, `value="http://www.example.com"`), ShouldNotEqual, -1)
			So(strings.Index(output, `maxlength="1000"`), ShouldNotEqual, -1)
			So(strings.Index(output, `for="urlFormHtml-url"`), ShouldNotEqual, -1)
			So(strings.Index(output, `>Enter URL:</label>`), ShouldNotEqual, -1)
			So(strings.Index(output, `<span class="help-block">I am UrlError</span>`), ShouldNotEqual, -1)
			So(strings.Index(output, `<span class="help-block">I am UrlError</span>`), ShouldNotEqual, -1)
			So(strings.Index(output, `class="form-group has-error"`), ShouldNotEqual, -1)

			testSubject.UrlErr = nil

			output = string(testSubject.GetHtml())

			So(strings.Index(output, `<span class="help-block">I am UrlError</span>`), ShouldEqual, -1)
			So(strings.Index(output, `class="form-group "`), ShouldNotEqual, -1)
		})

		Convey("UrlFieldValid", func() {
			testSubject.UrlModel = "http://www.example.com"
			So(testSubject.UrlFieldValid(), ShouldBeTrue)

			testSubject.UrlModel = "https://www.example.com"
			So(testSubject.UrlFieldValid(), ShouldBeTrue)

			testSubject.UrlModel = "I am not a valid url"
			So(testSubject.UrlFieldValid(), ShouldBeFalse)
		})
	})
}
