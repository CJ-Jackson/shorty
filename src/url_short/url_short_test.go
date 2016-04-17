package url_short

import (
	"fmt"
	"github.com/CJ-Jackson/shorty/src/http/http_error"
	"github.com/CJ-Jackson/shorty/src/http/http_redirect"
	"github.com/CJ-Jackson/shorty/src/random/random_mock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUrlShort(t *testing.T) {
	Convey("UrlShort", t, func() {
		chPanic := make(chan interface{})

		panicTrap := func(fn func()) {
			defer func() {
				chPanic <- recover()
			}()
			fn()
		}

		random := random_mock.NewRandomMock()
		repository := newUrlShortMongoRepositoryMock()

		testSubject := UrlShort{
			random:     random,
			repository: repository,
		}

		Convey("GetHash", func() {

			Convey("Manage to find hash for URL, therefore does not new hash", func(c C) {
				repository.C = c

				go func() {
					repository.ExpectFindByUrl("http://www.example.com", &urlShortData{
						Hash: "6bc01c99",
					}, nil)
				}()

				go panicTrap(func() {
					c.So(testSubject.GetHash("http://www.example.com"), ShouldEqual, "6bc01c99")
				})

				So(<-chPanic, ShouldBeNil)
			})

			Convey("Return err while trying to add data to the repository", func(c C) {
				random.C = c
				repository.C = c

				go func() {
					repository.ExpectFindByUrl("http://www.example.com", &urlShortData{},
						fmt.Errorf("I am error"))

					random.ExpectGenerateHex(URL_HASH_BYTE, "cc2fd4d0")

					repository.ExpectAddUrl(&urlShortData{
						Hash: "cc2fd4d0",
						Url:  "http://www.example.com",
					}, fmt.Errorf("Crap, I am out of order, sorry"))
				}()

				go panicTrap(func() {
					testSubject.GetHash("http://www.example.com")
				})

				So(<-chPanic, ShouldResemble,
					http_error.NewInternalServerErrorHttpHandler("UrlShort: unable to add url"))
			})

			Convey("Succefully add data to the repository and return hash nicely, with no error", func(c C) {
				random.C = c
				repository.C = c

				go func() {
					repository.ExpectFindByUrl("http://www.example.com", &urlShortData{},
						fmt.Errorf("I am error"))

					random.ExpectGenerateHex(URL_HASH_BYTE, "cc2fd4d0")

					repository.ExpectAddUrl(&urlShortData{
						Hash: "cc2fd4d0",
						Url:  "http://www.example.com",
					}, nil)
				}()

				go panicTrap(func() {
					c.So(testSubject.GetHash("http://www.example.com"), ShouldEqual, "cc2fd4d0")
				})

				So(<-chPanic, ShouldBeNil)
			})

		})

		Convey("DoRedirect", func() {

			Convey("Could not find url related to given hash, therefore Raise Not Found", func(c C) {
				repository.C = c

				go func() {
					repository.ExpectFindByHash("cc2fd4d0", &urlShortData{},
						fmt.Errorf("Could not find url."))
				}()

				go panicTrap(func() {
					testSubject.DoRedirect("cc2fd4d0")
				})

				So(<-chPanic, ShouldResemble,
					http_error.NewNotFoundErrorHttpHandler("UrlShort: didn't find url for hash"))
			})

			Convey("Found url related to given hash, therefore Raise Found", func(c C) {
				repository.C = c

				go func() {
					repository.ExpectFindByHash("cc2fd4d0", &urlShortData{
						Url: "http://www.example.com",
					}, nil)
				}()

				go panicTrap(func() {
					testSubject.DoRedirect("cc2fd4d0")
				})

				So(<-chPanic, ShouldResemble,
					http_redirect.NewFoundHttpHandler("http://www.example.com"))
			})

		})

	})
}
