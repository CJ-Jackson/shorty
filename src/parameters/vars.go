package parameters

import "sync"

var (
	shortyParameters = &ShortyParameters{}

	killSwitch = false
	killSwitchSync sync.Mutex
)
