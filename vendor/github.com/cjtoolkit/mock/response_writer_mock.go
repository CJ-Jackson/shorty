package mock

import (
	"bytes"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
)

type ResponseWriterMock struct {
	C C

	header http.Header

	write *bytes.Buffer

	writeHeaderParamStatus chan int
}

func NewResponseWriterMock() *ResponseWriterMock {
	return &ResponseWriterMock{
		header: http.Header{},
		write:  &bytes.Buffer{},
		writeHeaderParamStatus: make(chan int),
	}
}

func (w *ResponseWriterMock) Header() http.Header {
	return w.header
}

func (w *ResponseWriterMock) GetContent() string {
	defer w.write.Reset()
	return w.write.String()
}

func (w *ResponseWriterMock) Write(p []byte) (int, error) {
	return w.write.Write(p)
}

func (w *ResponseWriterMock) ExpectWriteHeader(expectStatus int) {
	w.writeHeaderParamStatus <- expectStatus
}

func (w *ResponseWriterMock) WriteHeader(status int) {
	w.C.So(status, ShouldEqual, <-w.writeHeaderParamStatus)
}
