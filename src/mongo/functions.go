package mongo

import (
	"github.com/CJ-Jackson/shorty/src/common"
	"github.com/CJ-Jackson/shorty/src/parameters"
	"gopkg.in/mgo.v2"
)

func InitShortyMongoDb() {
	killSwitchSync.Lock()
	defer killSwitchSync.Unlock()

	if killSwitch {
		return
	}
	killSwitch = true

	var err error
	session, err = mgo.Dial(parameters.GetShortyParameters().MongoDial)
	common.CheckError(err)
	session.SetMode(mgo.Monotonic, true)
}
