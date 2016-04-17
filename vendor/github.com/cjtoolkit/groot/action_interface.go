package groot

import "net/http"

type ActionInterface interface {
	Paths() []string
	New() http.Handler
}
