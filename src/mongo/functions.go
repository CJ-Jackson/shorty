package mongo

import (
	"github.com/CJ-Jackson/shorty/src/common"
	"github.com/CJ-Jackson/shorty/src/parameters"
	"gopkg.in/mgo.v2"
)

func InitShortyMongoDb() {
	var err error
	session, err = mgo.Dial(parameters.GetShortyParameters().MongoDial)
	session.SetMode(mgo.Monotonic, true)
	common.CheckError(err)
}
