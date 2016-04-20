package router

import (
	"github.com/cjtoolkit/groot"
	"sync"
)

var (
	router = groot.New()

	killSwitch = false
	killSwitchSync sync.Mutex
)
