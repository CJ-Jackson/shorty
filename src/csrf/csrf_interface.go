package csrf

import html "html/template"

type CsrfInterface interface {
	Token() string
	Field() html.HTML
}
