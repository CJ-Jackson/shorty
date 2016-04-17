package groot

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type HttpRouterInterface interface {
	Handle(method, path string, handle httprouter.Handle)
	Handler(method, path string, handler http.Handler)
	Lookup(method, path string) (httprouter.Handle, httprouter.Params, bool)
	ServeFiles(path string, root http.FileSystem)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}
