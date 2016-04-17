package groot

import "net/http"

type MethodPatchInterface interface {
	Patch()
}

func patch(handler http.Handler) {
	handler.(MethodPatchInterface).Patch()
}
