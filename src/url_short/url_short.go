package url_short

import (
	"github.com/CJ-Jackson/shorty/src/http/http_error"
	"github.com/CJ-Jackson/shorty/src/http/http_redirect"
	"github.com/CJ-Jackson/shorty/src/random"
)

type UrlShort struct {
	random     random.RandomInterface
	repository urlShortMongoRepositoryInterface
}

func (uS UrlShort) GetHash(urlStr string) string {
	data, err := uS.repository.FindByUrl(urlStr)
	if nil == err {
		return data.Hash
	}

	data.Hash = uS.random.GenerateHex(URL_HASH_BYTE)
	data.Url = urlStr

	err = uS.repository.AddUrl(data)
	if nil != err {
		http_error.RaiseInternalServerError("UrlShort: unable to add url")
	}

	return data.Hash
}

func (uS UrlShort) DoRedirect(hashStr string) {
	data, err := uS.repository.FindByHash(hashStr)
	if nil != err {
		http_error.RaiseNotFound("UrlShort: didn't find url for hash")
	}

	http_redirect.RaiseFound(data.Url)
}
