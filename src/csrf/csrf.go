package csrf

import (
	"github.com/gorilla/csrf"
	html "html/template"
	"net/http"
)

type Csrf struct {
	r *http.Request
}

func (c Csrf) Token() string {
	return csrf.Token(c.r)
}

func (c Csrf) Field() html.HTML {
	return csrf.TemplateField(c.r)
}
