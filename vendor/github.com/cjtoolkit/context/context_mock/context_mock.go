package context_mock

import (
	. "github.com/smartystreets/goconvey/convey"
)

type ContextMock struct {
	C C

	getParamName  chan string
	getWillReturn chan interface{}

	setParamName chan string
	setParamV    chan interface{}

	dealerParamName chan string

	expectAddFunc chan bool

	expectExecFuncs chan bool
}

func NewContextMock() *ContextMock {
	return &ContextMock{
		getParamName:  make(chan string),
		getWillReturn: make(chan interface{}),

		setParamName: make(chan string),
		setParamV:    make(chan interface{}),

		dealerParamName: make(chan string),

		expectAddFunc: make(chan bool),

		expectExecFuncs: make(chan bool),
	}
}

func (cM *ContextMock) ExpectGet(expectName string, willReturn interface{}) {
	cM.getParamName <- expectName
	cM.getWillReturn <- willReturn
}

func (cM *ContextMock) Get(name string) interface{} {
	So(name, ShouldEqual, <-cM.getParamName)
	return <-cM.getWillReturn
}

func (cM *ContextMock) ExpectSet(expectName string, willReturn interface{}) {
	cM.setParamName <- expectName
	cM.setParamV <- willReturn
}

func (cM *ContextMock) Set(name string, v interface{}) {
	So(name, ShouldEqual, <-cM.setParamName)
	So(name, ShouldResemble, <-cM.setParamV)
}

func (cM *ContextMock) ExpectDealer(expectName string) {
	cM.dealerParamName <- expectName
}

func (cM *ContextMock) Dealer(name string, gotIt func(v interface{}), dontHaveIt func() interface{}) {
	So(name, ShouldEqual, <-cM.dealerParamName)
}

func (cM *ContextMock) ExpectAddFunc() {
	cM.expectAddFunc <- true
}

func (cM *ContextMock) AddFunc(fn func()) {
	<-cM.expectAddFunc
}

func (cM *ContextMock) ExpectExecFuncs() {
	cM.expectExecFuncs <- true
}

func (cM *ContextMock) ExecFuncs() {
	<-cM.expectExecFuncs
}
