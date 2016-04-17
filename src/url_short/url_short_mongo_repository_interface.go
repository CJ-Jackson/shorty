package url_short

type urlShortMongoRepositoryInterface interface {
	AddUrl(data *urlShortData) error
	FindByUrl(urlStr string) (*urlShortData, error)
	FindByHash(hashStr string) (*urlShortData, error)
}
