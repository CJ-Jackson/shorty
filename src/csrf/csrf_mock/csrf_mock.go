package csrf_mock

import (
	. "github.com/smartystreets/goconvey/convey"
	html "html/template"
)

type CsrfMock struct {
	C C

	tokenWillReturn chan string

	fieldWillReturn chan string
}

func NewCsrfMock() *CsrfMock {
	return &CsrfMock{
		tokenWillReturn: make(chan string),
		fieldWillReturn: make(chan string),
	}
}

func (c *CsrfMock) ExpectToken(willReturn string) {
	c.tokenWillReturn <- willReturn
}

func (c *CsrfMock) Token() string {
	return <-c.tokenWillReturn
}

func (c *CsrfMock) ExpectField(willReturn string) {
	c.fieldWillReturn <- willReturn
}

func (c *CsrfMock) Field() html.HTML {
	return html.HTML(<-c.fieldWillReturn)
}
