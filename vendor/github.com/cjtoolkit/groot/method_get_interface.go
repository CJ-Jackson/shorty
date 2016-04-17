package groot

import "net/http"

type MethodGetInterface interface {
	Get()
}

func get(handler http.Handler) {
	handler.(MethodGetInterface).Get()
}
