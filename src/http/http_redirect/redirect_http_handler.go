package http_redirect

import "net/http"

type RedirectHttpHandler struct {
	Status int
	Url    string
}

func (rH RedirectHttpHandler) Error() string {
	return http.StatusText(rH.Status)
}

func (rH RedirectHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, rH.Url, rH.Status)
}

func NewMovedPermanentlyHttpHandler(urlStr string) RedirectHttpHandler {
	return RedirectHttpHandler{
		Status: http.StatusMovedPermanently,
		Url:    urlStr,
	}
}

func RaiseMovedPermanently(urlStr string) {
	panic(NewMovedPermanentlyHttpHandler(urlStr))
}

func NewFoundHttpHandler(urlStr string) RedirectHttpHandler {
	return RedirectHttpHandler{
		Status: http.StatusFound,
		Url:    urlStr,
	}
}

func RaiseFound(urlStr string) {
	panic(NewFoundHttpHandler(urlStr))
}

func NewSeeOtherHttpHandler(urlStr string) RedirectHttpHandler {
	return RedirectHttpHandler{
		Status: http.StatusSeeOther,
		Url:    urlStr,
	}
}

func RaiseSeeOther(urlStr string) {
	panic(NewSeeOtherHttpHandler(urlStr))
}
