package http_error

import (
	"fmt"
	"net/http"
)

type ErrorHandler struct {
	R interface{}
}

func (eH ErrorHandler) getHttpHandler() (http.Handler, int) {
	switch value := eH.R.(type) {
	case nil:
		// do nothing
	case http.Handler:
		return value, 1
	case error:
		return NewInternalServerErrorHttpHandler(value.Error()), 2
	default:
		return NewInternalServerErrorHttpHandler(fmt.Sprint(value)), 3
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}), 0
}

func (eH ErrorHandler) GetHttpHandler() http.Handler {
	h, _ := eH.getHttpHandler()
	return h
}
