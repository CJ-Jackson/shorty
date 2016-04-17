package groot

import "net/http"

type MethodPutInterface interface {
	Put()
}

func put(handler http.Handler) {
	handler.(MethodPutInterface).Put()
}
