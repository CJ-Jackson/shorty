package groot

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ParamsInterface interface {
	Params(params Params)
}

func params(handler http.Handler, params httprouter.Params) {
	handler.(ParamsInterface).Params(Params(params))
}

func blankParams(handler http.Handler, params httprouter.Params) {}
