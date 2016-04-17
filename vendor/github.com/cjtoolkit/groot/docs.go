/*
An Object Oriented Style router built on top of the speedy httprouter (https://github.com/julienschmidt/httprouter).

Usage

Setup a new router

	var router = groot.New()

Create a new action (Well a data type that implement ActionInterface)

	type indexAction struct {
		page string

		w http.ResponseWriter
		r *http.Request
	}

	// The principle is the same as https://godoc.org/github.com/julienschmidt/httprouter, because it's built on top
	// of that, it's an excellent foundation.
	func (ia indexAction) Paths() []string {
		return []string{
			"/",
			"/p/:page",
		}
	}

	// Great opportunity to set default values for url parameters.
	func (ia indexAction) New() http.Handler {
		return &indexAction{
			page: "1",
		}
	}

	// Optional Method: Populate field with paramters source from url. (This gets called first before ServeHTTP)
	func (ia *indexAction) Params(params groot.Params) {
		if page := params.ByName("page"); "" != page {
			ia.page = page
		}
	}

	// Good opportunity for dependency injection. (This gets called after Params and before HTTP Method.)
	// Even treat them http.ReponseWriter and *http.Request as dependencies.
	func (ia *indexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		ia.w, ia.r = w, r
	}

	// Groot requires at least one method that relate to http method, otherwise it will panic while getting
	// registered. There also Delete(), Head(), Options(), Patch(), Post() and Put(), non of them have parameters,
	// this is to encourage dependency injection.
	func (ia *indexAction) Get() {
		fmt.Fprintln(ia.w, "Hello World!")
		fmt.Fprintln(ia.w, ia.page)
	}

Than register the action.

	func init() {
		router.RegisterAction(indexAction{})
	}

Finally start the engine.

	func main() {
		http.ListenAndServe(":8080", router)
	}

*/
package groot
