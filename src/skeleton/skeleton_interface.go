package skeleton

import (
	"io"
)

type SkeletonInterface interface {
	SetStatus(status int)
	SetTitle(title string)
	SetHead(head io.Reader)
	SetBody(body io.Reader)
	SetFooter(footer io.Reader)
	SetJavascript(javascript io.Reader)
	Execute()
}
