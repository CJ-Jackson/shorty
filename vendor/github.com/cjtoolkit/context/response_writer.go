package context

import "net/http"

type responseWriter struct {
	http.ResponseWriter
	c *Context
}

func New(w http.ResponseWriter) (http.ResponseWriter, *Context) {
	res := &responseWriter{
		ResponseWriter: w,
		c:              newContext(),
	}
	return res, res.c
}

func GetOriginalResponseWriter(w http.ResponseWriter) http.ResponseWriter {
	return w.(responseWriter).ResponseWriter
}
