package skeleton

import (
	"github.com/CJ-Jackson/shorty/src/common"
	html "html/template"
	"io"
	"net/http"
)

type Skeleton struct {
	tpl        *html.Template
	w          http.ResponseWriter
	r          *http.Request
	status     int
	title      string
	head       io.Reader
	body       io.Reader
	footer     io.Reader
	javascript io.Reader
}

func (s *Skeleton) SetStatus(status int) {
	s.status = status
}

func (s *Skeleton) SetTitle(title string) {
	s.title = title
}

func (s *Skeleton) Title() string {
	return s.title
}

func (s *Skeleton) SetHead(head io.Reader) {
	s.head = head
}

func (s *Skeleton) Head() html.HTML {
	return common.ReaderToHtml(s.head)
}

func (s *Skeleton) SetBody(body io.Reader) {
	s.body = body
}

func (s *Skeleton) Body() html.HTML {
	return common.ReaderToHtml(s.body)
}

func (s *Skeleton) SetFooter(footer io.Reader) {
	s.footer = footer
}

func (s *Skeleton) Footer() html.HTML {
	return common.ReaderToHtml(s.footer)
}

func (s *Skeleton) SetJavascript(javascript io.Reader) {
	s.javascript = javascript
}

func (s *Skeleton) Javascript() html.HTML {
	return common.ReaderToHtml(s.javascript)
}

func (s *Skeleton) execute(wr io.Writer) {
	s.tpl.Execute(wr, s)
}

func (s *Skeleton) Execute() {
	s.w.Header().Set("Content-Type", "text/html; charset=utf-8")
	s.w.WriteHeader(s.status)
	s.execute(s.w)
}
