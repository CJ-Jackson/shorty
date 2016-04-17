package csrf

import (
	"github.com/CJ-Jackson/shorty/src/common"
	"github.com/cjtoolkit/context"
	"net/http"
)

func GetShortyCsrf(w http.ResponseWriter, r *http.Request) (c Csrf) {
	ctx := context.GetContext(w)
	ctx.Dealer(CSRF_CONTEXT_SERIAL, func(v interface{}) {
		c = v.(Csrf)
	}, func() interface{} {
		ch := make(chan interface{})
		go func() {
			defer func() {
				ch <- recover()
			}()
			csrfSystem(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ch <- nil
				<-ch
			})).ServeHTTP(w, r)
		}()
		common.PanicIfNotNil(<-ch)
		ctx.AddFunc(func() {
			ch <- nil
			<-ch
			close(ch)
		})
		c = Csrf{r}
		w.Header().Set("X-CSRF-Token", c.Token())
		return c
	})
	return
}
