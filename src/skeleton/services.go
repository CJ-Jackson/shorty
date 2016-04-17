package skeleton

import (
	"net/http"
)

func GetShortySkeleton(w http.ResponseWriter, r *http.Request) *Skeleton {
	return &Skeleton{
		tpl:    skeletonHtml,
		w:      w,
		r:      r,
		status: http.StatusOK,
	}
}
