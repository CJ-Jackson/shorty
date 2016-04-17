package skeleton_mock

import (
	"github.com/CJ-Jackson/shorty/src/common"
	. "github.com/smartystreets/goconvey/convey"
	html "html/template"
	"io"
)

type SkeletonMock struct {
	C C

	setStatusParamStatus chan int

	setTitleParamTitle chan string

	setHeadFetchReader chan io.Reader

	setBodyFetchReader chan io.Reader

	setFooterFetchReader chan io.Reader

	setJavascriptFetchReader chan io.Reader

	executeExpected chan bool
}

func NewSkeltonMock() *SkeletonMock {
	return &SkeletonMock{
		setStatusParamStatus: make(chan int),

		setTitleParamTitle: make(chan string),

		setHeadFetchReader: make(chan io.Reader),

		setBodyFetchReader: make(chan io.Reader),

		setFooterFetchReader: make(chan io.Reader),

		setJavascriptFetchReader: make(chan io.Reader),

		executeExpected: make(chan bool),
	}
}

func (sM *SkeletonMock) ExpectSetStatus(expectStatus int) {
	sM.setStatusParamStatus <- expectStatus
}

func (sM *SkeletonMock) SetStatus(status int) {
	sM.C.So(status, ShouldEqual, <-sM.setStatusParamStatus)
}

func (sM *SkeletonMock) ExpectSetTitle(expectTitle string) {
	sM.setTitleParamTitle <- expectTitle
}

func (sM *SkeletonMock) SetTitle(title string) {
	sM.C.So(title, ShouldEqual, <-sM.setTitleParamTitle)
}

func (sM *SkeletonMock) ExpectSetHead() html.HTML {
	return common.ReaderToHtml(<-sM.setHeadFetchReader)
}

func (sM *SkeletonMock) SetHead(head io.Reader) {
	sM.setHeadFetchReader <- head
}

func (sM *SkeletonMock) ExpectSetBody() html.HTML {
	return common.ReaderToHtml(<-sM.setBodyFetchReader)
}

func (sM *SkeletonMock) SetBody(body io.Reader) {
	sM.setBodyFetchReader <- body
}

func (sM *SkeletonMock) ExpectSetFooter() html.HTML {
	return common.ReaderToHtml(<-sM.setFooterFetchReader)
}

func (sM *SkeletonMock) SetFooter(footer io.Reader) {
	sM.setFooterFetchReader <- footer
}

func (sM *SkeletonMock) ExpectSetJavascript() html.HTML {
	return common.ReaderToHtml(<-sM.setJavascriptFetchReader)
}

func (sM *SkeletonMock) SetJavascript(javascript io.Reader) {
	sM.setJavascriptFetchReader <- javascript
}

func (sM *SkeletonMock) ExpectExecute() {
	sM.executeExpected <- true
}

func (sM *SkeletonMock) Execute() {
	<-sM.executeExpected
}
