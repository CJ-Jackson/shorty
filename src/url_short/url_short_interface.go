package url_short

type UrlShortInterface interface {
	GetHash(urlStr string) string
	DoRedirect(hashStr string)
}
