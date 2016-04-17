package groot

import "net/http"

type MethodDeleteInterface interface {
	Delete()
}

func delete(handler http.Handler) {
	handler.(MethodDeleteInterface).Delete()
}
