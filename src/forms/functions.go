package forms

import (
	html "html/template"
)

func ErrorToHtml(err error) html.HTML {
	if nil == err {
		return html.HTML("")
	}

	return html.HTML(`<span class="help-block">` + html.HTMLEscapeString(err.Error()) + `</span>`)
}
