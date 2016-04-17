package csrf

import "net/http"

var csrfSystem func(http.Handler) http.Handler
