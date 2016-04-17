package csrf

import (
	"github.com/CJ-Jackson/shorty/src/globals"
	"github.com/CJ-Jackson/shorty/src/http/http_error"
	"github.com/CJ-Jackson/shorty/src/parameters"
	"github.com/gorilla/csrf"
	"net/http"
)

func InitShortyCsrf() {
	csrfSystem = csrf.Protect(
		[]byte(parameters.GetShortyParameters().CsrfKey),
		csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http_error.RaiseForbidden("Invalid CSRF")
		})),
		csrf.Secure(globals.GetShortyGlobals().Production),
	)
}
