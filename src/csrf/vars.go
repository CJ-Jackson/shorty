package csrf

import (
	"net/http"
	"sync"
)

var (
	csrfSystem func(http.Handler) http.Handler

	killSwitch = false
	killSwitchSync sync.Mutex
)
