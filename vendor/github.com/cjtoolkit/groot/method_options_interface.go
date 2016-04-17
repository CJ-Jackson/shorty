package groot

import "net/http"

type MethodOptionsInterface interface {
	Options()
}

func options(handler http.Handler) {
	handler.(MethodOptionsInterface).Options()
}
