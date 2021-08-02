package wkwsweb

import "net/http"

type Routers []Router

type Router struct {
	Handler http.HandlerFunc
	Path    string
	Method  string
}

func (wkws *Wkws) POST(uri string, handler http.HandlerFunc) {
	AddRouter(http.MethodPost, uri, handler, wkws)
}

func (wkws *Wkws) GET(uri string, handler http.HandlerFunc) {
	AddRouter(http.MethodGet, uri, handler, wkws)
}

func AddRouter(method string, path string, handler http.HandlerFunc, core *Wkws) {
	core.RouterGroup = append(core.RouterGroup, Router{
		Handler: handler,
		Path:    path,
		Method:  method,
	})
}
