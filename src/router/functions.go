package router

import (
	"github.com/CJ-Jackson/shorty/src/http/http_error"
	"github.com/CJ-Jackson/shorty/src/parameters"
	"github.com/cjtoolkit/groot"
	"net/http"
)

func init() {
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http_error.RaiseNotFound("Router")
	})

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http_error.RaiseMethodNotAllowed("Router")
	})
}

func SetUpShortyFileServers() {
	killSwitchSync.Lock()
	defer killSwitchSync.Unlock()

	if killSwitch {
		return
	}
	killSwitch = true

	router.Router.ServeFiles("/static/*filepath", http.Dir(parameters.GetShortyParameters().FilePath))
}

func RegisterShortyAction(action groot.ActionInterface) {
	router.RegisterAction(action)
}

func ServeShortyRouterHTTP(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
