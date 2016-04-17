package url_short

import (
	. "github.com/smartystreets/goconvey/convey"
)

type urlShortMongoRepositoryMock struct {
	C C

	addUrlParamData  chan *urlShortData
	addUrlWillReturn chan error

	findByUrlParamUrlStr            chan string
	findByUrlWillReturnUrlShortData chan *urlShortData
	findByUrlWillReturnError        chan error

	findByHashParamHashStr           chan string
	findByHashWillReturnUrlShortData chan *urlShortData
	findByHashWillReturnError        chan error
}

func newUrlShortMongoRepositoryMock() *urlShortMongoRepositoryMock {
	return &urlShortMongoRepositoryMock{
		addUrlParamData:  make(chan *urlShortData),
		addUrlWillReturn: make(chan error),

		findByUrlParamUrlStr:            make(chan string),
		findByUrlWillReturnUrlShortData: make(chan *urlShortData),
		findByUrlWillReturnError:        make(chan error),

		findByHashParamHashStr:           make(chan string),
		findByHashWillReturnUrlShortData: make(chan *urlShortData),
		findByHashWillReturnError:        make(chan error),
	}
}

func (uSMR *urlShortMongoRepositoryMock) ExpectAddUrl(expectData *urlShortData, willReturn error) {
	uSMR.addUrlParamData <- expectData

	uSMR.addUrlWillReturn <- willReturn
}

func (uSMR *urlShortMongoRepositoryMock) AddUrl(data *urlShortData) error {
	uSMR.C.So(data, ShouldResemble, <-uSMR.addUrlParamData)

	return <-uSMR.addUrlWillReturn
}

func (uSMR *urlShortMongoRepositoryMock) ExpectFindByUrl(
	expectUrlStr string,
	willReturnUrlShortData *urlShortData,
	willReturnError error,
) {
	uSMR.findByUrlParamUrlStr <- expectUrlStr

	uSMR.findByUrlWillReturnUrlShortData <- willReturnUrlShortData
	uSMR.findByUrlWillReturnError <- willReturnError
}

func (uSMR *urlShortMongoRepositoryMock) FindByUrl(urlStr string) (*urlShortData, error) {
	uSMR.C.So(urlStr, ShouldEqual, <-uSMR.findByUrlParamUrlStr)

	return <-uSMR.findByUrlWillReturnUrlShortData, <-uSMR.findByUrlWillReturnError
}

func (uSMR *urlShortMongoRepositoryMock) ExpectFindByHash(
	expectHashStr string,
	willReturnUrlShortData *urlShortData,
	willReturnError error,
) {
	uSMR.findByHashParamHashStr <- expectHashStr

	uSMR.findByHashWillReturnUrlShortData <- willReturnUrlShortData
	uSMR.findByHashWillReturnError <- willReturnError
}

func (uSMR *urlShortMongoRepositoryMock) FindByHash(hashStr string) (*urlShortData, error) {
	uSMR.C.So(hashStr, ShouldEqual, <-uSMR.findByHashParamHashStr)

	return <-uSMR.findByHashWillReturnUrlShortData, <-uSMR.findByHashWillReturnError
}
