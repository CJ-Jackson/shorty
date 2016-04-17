package web

import (
	"github.com/CJ-Jackson/shorty/src/csrf/csrf_mock"
	"github.com/CJ-Jackson/shorty/src/forms"
	"github.com/CJ-Jackson/shorty/src/skeleton/skeleton_mock"
	"github.com/CJ-Jackson/shorty/src/url_short/url_short_mock"
	"github.com/cjtoolkit/form/form_mock"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestIndexAction(t *testing.T) {
	Convey("indexAction", t, func() {
		r := &http.Request{
			Form: url.Values{},
		}

		formMock := form_mock.NewFormMock()

		sMock := skeleton_mock.NewSkeltonMock()

		urlShortMock := url_short_mock.NewUrlShortMock()

		csrfMock := csrf_mock.NewCsrfMock()

		testSubject := &indexAction{
			r:        r,
			headTpl:  strings.NewReader(indexActionHeadHtml),
			bodyTpl:  indexActionBodyHtml,
			UrlForm:  forms.NewUrlForm(),
			form:     formMock,
			S:        sMock,
			urlShort: urlShortMock,
			Csrf:     csrfMock,
		}

		Convey("Form failed to validate", func(c C) {
			formMock.C = c
			sMock.C = c
			urlShortMock.C = c

			go func() {
				formMock.ExpectSetForm(r.Form)
				formMock.ExpectValidate(testSubject.UrlForm, false)

				formMock.ExpectTransform(testSubject.UrlForm, false)
				testSubject.UrlForm.Fields()

				csrfMock.ExpectField(`<!-- TokenField -->`)

				sMock.ExpectSetTitle("Welcome")

				bodyStr := string(sMock.ExpectSetBody())

				c.So(strings.Index(bodyStr,
					`<div class="alert alert-danger" role="alert">Hmm... something did not go right!</div>`),
					ShouldNotEqual, -1)

				c.So(strings.Index(bodyStr,
					`<!-- TokenField -->`),
					ShouldNotEqual, -1)

				c.So(strings.Index(bodyStr,
					string(testSubject.UrlForm.GetHtml())),
					ShouldNotEqual, -1)

				sMock.ExpectSetHead()
				sMock.ExpectExecute()
			}()

			testSubject.Post()
		})

		Convey("Url not valid in form", func(c C) {
			formMock.C = c
			sMock.C = c
			urlShortMock.C = c

			testSubject.UrlForm.UrlModel = "aaaaaaaaaa"

			go func() {
				formMock.ExpectSetForm(r.Form)
				formMock.ExpectValidate(testSubject.UrlForm, true)

				formMock.ExpectTransform(testSubject.UrlForm, false)
				testSubject.UrlForm.Fields()

				csrfMock.ExpectField(`<!-- TokenField -->`)

				sMock.ExpectSetTitle("Welcome")

				bodyStr := string(sMock.ExpectSetBody())

				c.So(strings.Index(bodyStr,
					`<div class="alert alert-danger" role="alert">Hmm... something did not go right!</div>`),
					ShouldNotEqual, -1)

				c.So(strings.Index(bodyStr,
					`<!-- TokenField -->`),
					ShouldNotEqual, -1)

				c.So(strings.Index(bodyStr,
					string(testSubject.UrlForm.GetHtml())),
					ShouldNotEqual, -1)

				sMock.ExpectSetHead()
				sMock.ExpectExecute()
			}()

			testSubject.Post()
		})

		Convey("Valid Form (http://)", func(c C) {
			formMock.C = c
			sMock.C = c
			urlShortMock.C = c

			testSubject.UrlForm.UrlModel = "http://"

			go func() {
				formMock.ExpectSetForm(r.Form)
				formMock.ExpectValidate(testSubject.UrlForm, true)

				urlShortMock.ExpectGetHash("http://", "abc123")

				formMock.ExpectTransform(testSubject.UrlForm, false)
				testSubject.UrlForm.Fields()

				csrfMock.ExpectField(`<!-- TokenField -->`)

				sMock.ExpectSetTitle("Welcome")

				bodyStr := string(sMock.ExpectSetBody())

				c.So(strings.Index(bodyStr,
					`<div class="alert alert-success" role="alert">Your new url is <a href="/r/abc123">/r/abc123</a>!</div>`),
					ShouldNotEqual, -1)

				c.So(strings.Index(bodyStr,
					`<!-- TokenField -->`),
					ShouldNotEqual, -1)

				c.So(strings.Index(bodyStr,
					string(testSubject.UrlForm.GetHtml())),
					ShouldNotEqual, -1)

				sMock.ExpectSetHead()
				sMock.ExpectExecute()
			}()

			testSubject.Post()
		})

		Convey("Valid Form (https://)", func(c C) {
			formMock.C = c
			sMock.C = c
			urlShortMock.C = c

			testSubject.UrlForm.UrlModel = "https://"

			go func() {
				formMock.ExpectSetForm(r.Form)
				formMock.ExpectValidate(testSubject.UrlForm, true)

				urlShortMock.ExpectGetHash("https://", "abc123")

				formMock.ExpectTransform(testSubject.UrlForm, false)
				testSubject.UrlForm.Fields()

				csrfMock.ExpectField(`<!-- TokenField -->`)

				sMock.ExpectSetTitle("Welcome")

				bodyStr := string(sMock.ExpectSetBody())

				c.So(strings.Index(bodyStr,
					`<div class="alert alert-success" role="alert">Your new url is <a href="/r/abc123">/r/abc123</a>!</div>`),
					ShouldNotEqual, -1)

				c.So(strings.Index(bodyStr,
					`<!-- TokenField -->`),
					ShouldNotEqual, -1)

				c.So(strings.Index(bodyStr,
					string(testSubject.UrlForm.GetHtml())),
					ShouldNotEqual, -1)

				sMock.ExpectSetHead()
				sMock.ExpectExecute()
			}()

			testSubject.Post()
		})
	})
}
