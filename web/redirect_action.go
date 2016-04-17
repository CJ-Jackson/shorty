package web

import (
	"github.com/CJ-Jackson/shorty/src/router"
	"github.com/CJ-Jackson/shorty/src/url_short"
	"github.com/cjtoolkit/groot"
	"net/http"
)

type redirectAction struct {
	hash string

	urlShort url_short.UrlShortInterface
}

func (rA redirectAction) Paths() []string {
	return []string{
		"/r/:hash",
	}
}

func (rA redirectAction) New() http.Handler {
	return &redirectAction{}
}

func (rA *redirectAction) Params(params groot.Params) {
	rA.hash = params.ByName("hash")
}

func (rA *redirectAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rA.urlShort = url_short.GetUrlShort()
}

func (rA *redirectAction) Get() {
	rA.urlShort.DoRedirect(rA.hash)
}

func init() {
	router.RegisterShortyAction(redirectAction{})
}
