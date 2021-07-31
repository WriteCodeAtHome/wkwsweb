package wkwsweb

import "net/http"

func (wkws *Wkws) POST(uri string, handler http.HandlerFunc) {
	wkws.Router.HandleFunc(uri, handler)
}

func (wkws *Wkws) GET(uri string, handler http.HandlerFunc) {
	wkws.Router.HandleFunc(uri, handler)
}
