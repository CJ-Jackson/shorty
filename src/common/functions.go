package common

import (
	html "html/template"
	"io"
	"io/ioutil"
)

func CheckError(err error) {
	if nil != err {
		panic(err)
	}
}

func PanicIfNotNil(v interface{}) {
	if nil != v {
		panic(v)
	}
}

func renderToHtml(reader io.Reader) (h html.HTML, i int) {
	if nil == reader {
		i = 1
		return
	}

	b, err := ioutil.ReadAll(reader)
	if nil != err {
		i = 2
		return
	}

	h = html.HTML(b)
	return
}

func ReaderToHtml(reader io.Reader) (h html.HTML) {
	h, _ = renderToHtml(reader)
	return
}

func TrueFalseExecFunction(b bool, trueFn func(), falseFn func()) {
	if b {
		trueFn()
	} else {
		falseFn()
	}
}

func ExecFunctionIfTrue(b bool, trueFn func()) {
	TrueFalseExecFunction(b, trueFn, func() {})
}

func ExecFunctionIfFalse(b bool, falseFn func()) {
	TrueFalseExecFunction(b, func() {}, falseFn)
}
