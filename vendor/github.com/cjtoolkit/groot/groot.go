package groot

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// It's is recommended to call New() rather than manually construct Groot, you can always change the field after,
// I advise you to leave Router field alone. And don't set NotFound and MethodNotAllowed to nil, it will panic.
type Groot struct {
	Router           HttpRouterInterface
	NotFound         http.Handler
	MethodNotAllowed http.Handler
}

func New() *Groot {
	return &Groot{
		Router: &httprouter.Router{
			RedirectTrailingSlash:  true,
			RedirectFixedPath:      true,
			HandleMethodNotAllowed: true,
			HandleOPTIONS:          true,
			NotFound: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				panic(errorNotFound{})
			}),
			MethodNotAllowed: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				panic(errorMethodNotAllowed{})
			}),
		},
		NotFound: http.NotFoundHandler(),
		MethodNotAllowed: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w,
				http.StatusText(http.StatusMethodNotAllowed),
				http.StatusMethodNotAllowed,
			)
		}),
	}
}

func (g *Groot) handleError(w http.ResponseWriter, r *http.Request) {
	switch recv := recover().(type) {
	case nil:
		// Do nothing.
	case errorNotFound:
		g.NotFound.ServeHTTP(w, r)
	case errorMethodNotAllowed:
		g.MethodNotAllowed.ServeHTTP(w, r)
	default:
		panic(recv)
	}
}

func (g *Groot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer g.handleError(w, r)
	g.Router.ServeHTTP(w, r)
}

// Not only the data type need to implement ActionInterface, it's also needs to to implement http.Handler for New()
// Method and it's need at least one method that relate to http method
// (e.g. Delete(), Get(), Head(), Options(), Patch(), Post(), Put())
func (g *Groot) RegisterAction(action ActionInterface) {
	g.panicIfActionIsNil(action)

	paramsFn := blankParams

	methods := methodMap{}

	handle := httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		handler := action.New()
		paramsFn(handler, params)
		handler.ServeHTTP(w, r)
		methods[r.Method](handler)
	})

	handler := action.New()
	g.panicIfHandlerIsNil(handler)

	g.checkAndSetupDelete(methods, handler)
	g.checkAndSetupGet(methods, handler)
	g.checkAndSetupHead(methods, handler)
	g.checkAndSetupOptions(methods, handler)
	g.checkAndSetupPatch(methods, handler)
	g.checkAndSetupPost(methods, handler)
	g.checkAndSetupPut(methods, handler)

	g.panicIfMethodHasNotBeenSpecified(methods)

	g.checkAndSetUpParams(handler, &paramsFn)

	paths := action.Paths()
	g.panicIfPathsIsNilOrEmpty(paths)

	for method := range methods {
		for _, path := range paths {
			g.Router.Handle(method, path, handle)
		}
	}
}

func (g *Groot) panicIfHandlerIsNil(handler http.Handler) {
	if nil == handler {
		panic("'handler' cannot be 'nil'")
	}
}

func (g *Groot) panicIfActionIsNil(action ActionInterface) {
	if nil == action {
		panic("'action' cannot be 'nil'")
	}
}

func (g *Groot) panicIfPathsIsNilOrEmpty(paths []string) {
	if 0 == len(paths) {
		panic("'paths' cannot be 'nil' or 'empty'")
	}
}

func (g *Groot) panicIfMethodHasNotBeenSpecified(methods methodMap) {
	if 0 == len(methods) {
		panic("A HTTP method has not been specified.")
	}
}

func (g *Groot) checkAndSetupDelete(methods methodMap, handler http.Handler) (b bool) {
	if _, ok := handler.(MethodDeleteInterface); ok {
		methods["DELETE"] = delete
		b = ok
	}
	return
}

func (g *Groot) checkAndSetupGet(methods methodMap, handler http.Handler) (b bool) {
	if _, ok := handler.(MethodGetInterface); ok {
		methods["GET"] = get
		b = ok
	}
	return
}

func (g *Groot) checkAndSetupHead(methods methodMap, handler http.Handler) (b bool) {
	if _, ok := handler.(MethodHeadInterface); ok {
		methods["HEAD"] = head
		b = ok
	}
	return
}

func (g *Groot) checkAndSetupOptions(methods methodMap, handler http.Handler) (b bool) {
	if _, ok := handler.(MethodOptionsInterface); ok {
		methods["OPTIONS"] = options
		b = ok
	}
	return
}

func (g *Groot) checkAndSetupPatch(methods methodMap, handler http.Handler) (b bool) {
	if _, ok := handler.(MethodPatchInterface); ok {
		methods["PATCH"] = patch
		b = ok
	}
	return
}

func (g *Groot) checkAndSetupPost(methods methodMap, handler http.Handler) (b bool) {
	if _, ok := handler.(MethodPostInterface); ok {
		methods["POST"] = post
		b = ok
	}
	return
}

func (g *Groot) checkAndSetupPut(methods methodMap, handler http.Handler) (b bool) {
	if _, ok := handler.(MethodPutInterface); ok {
		methods["PUT"] = put
		b = ok
	}
	return
}

func (g *Groot) checkAndSetUpParams(handler http.Handler,
	paramsFn *func(handler http.Handler, params httprouter.Params)) (b bool) {
	if _, ok := handler.(ParamsInterface); ok {
		*paramsFn = params
		b = ok
	}
	return
}
