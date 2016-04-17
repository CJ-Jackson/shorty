package url_short

import (
	"github.com/CJ-Jackson/shorty/src/mongo"
	"github.com/CJ-Jackson/shorty/src/random"
)

func getUrlShortMongoRepository() urlShortMongoRepository {
	return urlShortMongoRepository{
		collection: mongo.GetShortyMongoDatabase().C(URL_SHORT_MONGO_COLLECTION),
	}
}

func GetUrlShort() UrlShort {
	return UrlShort{
		random:     random.GetShortyRandom(),
		repository: getUrlShortMongoRepository(),
	}
}
