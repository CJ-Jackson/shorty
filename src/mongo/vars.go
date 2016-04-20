package mongo

import (
	"gopkg.in/mgo.v2"
	"sync"
)

var (
	session *mgo.Session

	killSwitch = false
	killSwitchSync sync.Mutex
)
