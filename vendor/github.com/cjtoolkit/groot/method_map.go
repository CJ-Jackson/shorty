package groot

import "net/http"

type methodMap map[string]func(handler http.Handler)
