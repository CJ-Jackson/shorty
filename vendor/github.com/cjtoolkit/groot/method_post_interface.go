package groot

import "net/http"

type MethodPostInterface interface {
	Post()
}

func post(handler http.Handler) {
	handler.(MethodPostInterface).Post()
}
