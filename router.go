package wkwsweb

import "net/http"

type Routers []Router

type Router struct {
	Handler Controller
	Path    string
	Method  string
}

func (wkws *Wkws) POST(uri string, handler Controller) {
	AddRouter(http.MethodPost, uri, handler, wkws)
}

func (wkws *Wkws) GET(uri string, handler Controller) {
	AddRouter(http.MethodGet, uri, handler, wkws)
}

func AddRouter(method string, path string, handler Controller, core *Wkws) {
	core.RouterGroup = append(core.RouterGroup, Router{
		Handler: handler,
		Path:    path,
		Method:  method,
	})
}
