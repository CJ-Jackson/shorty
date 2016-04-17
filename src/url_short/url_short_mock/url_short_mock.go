package url_short_mock

import . "github.com/smartystreets/goconvey/convey"

type UrlShortMock struct {
	C C

	getHashParamUrlStr chan string
	getHashWillReturn  chan string

	doRedirectParamHashStr chan string
}

func NewUrlShortMock() *UrlShortMock {
	return &UrlShortMock{
		getHashParamUrlStr: make(chan string),
		getHashWillReturn:  make(chan string),

		doRedirectParamHashStr: make(chan string),
	}
}

func (u *UrlShortMock) ExpectGetHash(expectUrlStr, willReturn string) {
	u.getHashParamUrlStr <- expectUrlStr
	u.getHashWillReturn <- willReturn
}

func (u *UrlShortMock) GetHash(urlStr string) string {
	u.C.So(urlStr, ShouldEqual, <-u.getHashParamUrlStr)

	return <-u.getHashWillReturn
}

func (u *UrlShortMock) ExpectDoRedirect(expectHashStr string) {
	u.doRedirectParamHashStr <- expectHashStr
}

func (u *UrlShortMock) DoRedirect(hashStr string) {
	u.C.So(hashStr, ShouldEqual, <-u.doRedirectParamHashStr)
}
