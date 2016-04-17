package mongo

import (
	"github.com/CJ-Jackson/shorty/src/parameters"
	"gopkg.in/mgo.v2"
)

func GetShortyMongoDatabase() *mgo.Database {
	return session.DB(parameters.GetShortyParameters().MongoDbName)
}
