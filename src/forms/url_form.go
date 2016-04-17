package forms

import (
	"bytes"
	"fmt"
	"github.com/cjtoolkit/form"
	"github.com/cjtoolkit/form/fields"
	html "html/template"
	"strings"
)

type UrlForm struct {
	fields []form.FormFieldInterface
	tpl    *html.Template

	UrlModel string
	UrlErr   error
	UrlNorm  string
}

func NewUrlForm() *UrlForm {
	return &UrlForm{
		tpl: urlFormHtml,
	}
}

func (u *UrlForm) Fields() []form.FormFieldInterface {
	u.fields = u.createFields()
	return u.fields
}

func (u *UrlForm) createFields() []form.FormFieldInterface {
	return []form.FormFieldInterface{
		fields.Text{
			Label:   "Enter URL",
			Name:    "url",
			Model:   &u.UrlModel,
			Err:     &u.UrlErr,
			Norm:    &u.UrlNorm,
			MinChar: 10,
			MaxChar: 1000,
		},
	}
}

func (u *UrlForm) HasUrlErr() bool {
	return nil != u.UrlErr
}

func (u *UrlForm) UrlErrHtml() html.HTML {
	return ErrorToHtml(u.UrlErr)
}

func (u *UrlForm) UrlField() fields.Text {
	return u.fields[0].(fields.Text)
}

func (u *UrlForm) UrlFieldValid() bool {
	value := strings.ToLower(u.UrlModel)
	if "http://" == value[:7] || "https://" == value[:8] {
		return true
	}
	u.UrlErr = fmt.Errorf("'%s' is not a valid url", u.UrlModel)
	return false
}

func (u *UrlForm) GetHtml() html.HTML {
	buf := &bytes.Buffer{}
	defer buf.Reset()
	u.tpl.Execute(buf, u)

	return html.HTML(buf.String())
}

func (u *UrlForm) Clear() {
	*u = *(NewUrlForm())
}
