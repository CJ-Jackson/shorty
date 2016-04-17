package http_error

import (
	"fmt"
	"github.com/CJ-Jackson/shorty/src/common"
	"github.com/CJ-Jackson/shorty/src/globals"
	"github.com/CJ-Jackson/shorty/src/skeleton"
	"io"
	"net/http"
	"runtime/debug"
	"strings"
)

type ErrorHttpHandler struct {
	Status   int
	ErrorMsg string
	Msg      string
}

func (e *ErrorHttpHandler) Error() string {
	return fmt.Sprintf("%d: %s: %s", e.Status, e.Error, e.Msg)
}

func (e *ErrorHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	s := skeleton.GetShortySkeleton(w, r)
	s.SetTitle(fmt.Sprintf("%d: %s", e.Status, e.ErrorMsg))
	s.SetBody(e.body())
	s.SetStatus(e.Status)
	s.Execute()
}

func (e *ErrorHttpHandler) body() (reader io.Reader) {
	common.TrueFalseExecFunction(globals.GetShortyGlobals().Production, func() {
		reader = strings.NewReader(fmt.Sprintf(`<h1>%d: %s</h1>`,
			e.Status, e.ErrorMsg))
	}, func() {
		reader = strings.NewReader(fmt.Sprintf(`<h1>%d: %s</h1><h2>%s</h2><pre>%s</pre>`,
			e.Status, e.ErrorMsg, e.Msg, debug.Stack()))
	})

	return
}

func NewNotFoundErrorHttpHandler(msg string) *ErrorHttpHandler {
	return &ErrorHttpHandler{
		Status:   http.StatusNotFound,
		ErrorMsg: http.StatusText(http.StatusNotFound),
		Msg:      msg,
	}
}

func RaiseNotFound(msg string) {
	panic(NewNotFoundErrorHttpHandler(msg))
}

func NewMethodNotAllowedHttpHandler(msg string) *ErrorHttpHandler {
	return &ErrorHttpHandler{
		Status:   http.StatusMethodNotAllowed,
		ErrorMsg: http.StatusText(http.StatusMethodNotAllowed),
		Msg:      msg,
	}
}

func RaiseMethodNotAllowed(msg string) {
	panic(NewMethodNotAllowedHttpHandler(msg))
}

func NewForbiddenHttpHandler(msg string) *ErrorHttpHandler {
	return &ErrorHttpHandler{
		Status:   http.StatusForbidden,
		ErrorMsg: http.StatusText(http.StatusForbidden),
		Msg:      msg,
	}
}

func RaiseForbidden(msg string) {
	panic(NewForbiddenHttpHandler(msg))
}

func NewInternalServerErrorHttpHandler(msg string) *ErrorHttpHandler {
	return &ErrorHttpHandler{
		Status:   http.StatusInternalServerError,
		ErrorMsg: http.StatusText(http.StatusInternalServerError),
		Msg:      msg,
	}
}

func RaiseInternalServerError(msg string) {
	panic(NewInternalServerErrorHttpHandler(msg))
}
