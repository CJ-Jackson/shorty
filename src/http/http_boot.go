package http

import (
	"fmt"
	"github.com/CJ-Jackson/shorty/src/http/http_error"
	"github.com/CJ-Jackson/shorty/src/router"
	"github.com/cjtoolkit/context"
	"net/http"
	"os"
)

type httpBoot struct {
	debug bool
}

func (hB httpBoot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w, c := context.New(w)

	defer func() {
		(http_error.ErrorHandler{R: recover()}).GetHttpHandler().ServeHTTP(w, r)
		c.ExecFuncs()
	}()

	hB.callFnOnDebugOnly(func() {
		fmt.Fprintln(os.Stderr, r.Method, r.URL.String())
	})

	router.ServeShortyRouterHTTP(w, r)
}

func (hB httpBoot) callFnOnDebugOnly(fn func()) {
	if hB.debug {
		fn()
	}
}
