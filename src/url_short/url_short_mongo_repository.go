package url_short

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type urlShortMongoRepository struct {
	collection *mgo.Collection
}

func (uSMR urlShortMongoRepository) AddUrl(data *urlShortData) error {
	data.CreatedAt = time.Now()
	return uSMR.collection.Insert(data)
}

func (uSMR urlShortMongoRepository) FindByUrl(urlStr string) (*urlShortData, error) {
	result := &urlShortData{}
	err := uSMR.collection.Find(bson.M{"url": urlStr}).One(result)

	return result, err
}

func (uSMR urlShortMongoRepository) FindByHash(hashStr string) (*urlShortData, error) {
	result := &urlShortData{}
	err := uSMR.collection.Find(bson.M{"hash": hashStr}).One(result)

	return result, err
}
