package groot

import "net/http"

type MethodHeadInterface interface {
	Head()
}

func head(handler http.Handler) {
	handler.(MethodHeadInterface).Head()
}
