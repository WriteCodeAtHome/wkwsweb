package wkwsweb

import (
	"net/http"
)

type RouterHandlers []RouterHandler

type RouterHandler struct {
	Path    string
	Method  string
	Handler Controller
}

func (wkws *Wkws) POST(uri string, handler Controller) {
	addRouter(http.MethodPost, uri, handler, wkws)
}

func (wkws *Wkws) GET(uri string, handler Controller) {
	addRouter(http.MethodGet, uri, handler, wkws)
}

func addRouter(method string, path string, handler Controller, core *Wkws) {
	core.RouterHandlers = append(core.RouterHandlers, RouterHandler{
		Path:    path,
		Method:  method,
		Handler: handler,
	})
}
