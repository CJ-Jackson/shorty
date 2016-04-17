package random_mock

import (
	. "github.com/smartystreets/goconvey/convey"
)

type RandomMock struct {
	C C

	generateHexParamNumOfBytes chan int
	generateHexWillReturn      chan string
}

func NewRandomMock() *RandomMock {
	return &RandomMock{
		generateHexParamNumOfBytes: make(chan int),
		generateHexWillReturn:      make(chan string),
	}
}

func (rand *RandomMock) ExpectGenerateHex(expectNumOfBytes int, willReturn string) {
	rand.generateHexParamNumOfBytes <- expectNumOfBytes
	rand.generateHexWillReturn <- willReturn
}

func (rand *RandomMock) GenerateHex(numOfBytes int) string {
	rand.C.So(numOfBytes, ShouldEqual, <-rand.generateHexParamNumOfBytes)

	return <-rand.generateHexWillReturn
}
