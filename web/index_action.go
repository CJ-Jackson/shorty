package web

import (
	"bytes"
	"fmt"
	"github.com/CJ-Jackson/shorty/src/csrf"
	"github.com/CJ-Jackson/shorty/src/forms"
	"github.com/CJ-Jackson/shorty/src/parameters"
	"github.com/CJ-Jackson/shorty/src/router"
	"github.com/CJ-Jackson/shorty/src/skeleton"
	"github.com/CJ-Jackson/shorty/src/url_short"
	"github.com/cjtoolkit/form"
	html "html/template"
	"io"
	"net/http"
	"strings"
)

type indexAction struct {
	w http.ResponseWriter
	r *http.Request

	Csrf     csrf.CsrfInterface
	headTpl  io.Reader
	bodyTpl  *html.Template
	UrlForm  *forms.UrlForm
	form     form.FormInterface
	S        skeleton.SkeletonInterface
	urlShort url_short.UrlShortInterface

	flash  string
	domain string
}

func (iA indexAction) Paths() []string {
	return []string{
		"/",
	}
}

func (iA indexAction) New() http.Handler {
	return &indexAction{}
}

func (iA *indexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	iA.w, iA.r = w, r
	iA.Csrf = csrf.GetShortyCsrf(w, r)
	iA.headTpl = strings.NewReader(indexActionHeadHtml)
	iA.bodyTpl = indexActionBodyHtml
	iA.UrlForm = forms.NewUrlForm()
	iA.form = form.NewFormDefaultLanguage()
	iA.S = skeleton.GetShortySkeleton(w, r)
	iA.urlShort = url_short.GetUrlShort()
	iA.domain = parameters.GetShortyParameters().Domain
}

func (iA *indexAction) Flash() html.HTML {
	return html.HTML(iA.flash)
}

func (iA *indexAction) Get() {
	iA.form.Transform(iA.UrlForm)

	body := &bytes.Buffer{}
	iA.bodyTpl.Execute(body, iA)

	iA.S.SetTitle("Welcome")
	iA.S.SetBody(body)
	iA.S.SetHead(iA.headTpl)
	iA.S.Execute()
}

func (iA *indexAction) Post() {
	iA.r.ParseForm()
	iA.form.SetForm(iA.r.Form)

	if !iA.form.Validate(iA.UrlForm) || !iA.UrlForm.UrlFieldValid() {
		iA.flash = `<div class="alert alert-danger" role="alert">Hmm... something did not go right!</div>`
		iA.Get()
		return
	}

	hash := iA.urlShort.GetHash(iA.UrlForm.UrlModel)

	iA.flash = fmt.Sprintf(
		`<div class="alert alert-success" role="alert">Your new url is <a href="/r/%s">%s/r/%s</a>!</div>`,
		hash, iA.domain, hash)

	iA.UrlForm.Clear()
	iA.Get()
}

func init() {
	router.RegisterShortyAction(indexAction{})
}
