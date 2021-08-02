package wkwsweb

import (
	"net/http"
)

type Error struct {
	Err error
}

type Wkws struct {
	RouterGroup Routers
}

func Init() (core *Wkws) {
	core = &Wkws{
		RouterGroup: nil,
	}
	return
}

func (wkws *Wkws) Run(add string, port string) (err error) {
	addr := add + ":" + port
	CLogger("HTTP listen the address , %s \n", addr)
	r := http.NewServeMux()
	CLogger("Router map:")
	for _, router := range wkws.RouterGroup {
		r.HandleFunc(router.Path, wkws.ServeHTTP)
		CLogger(router.Method + " " + router.Path)
	}
	err = http.ListenAndServe(addr, r)
	return
}

func (wkws *Wkws) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	c := NewCtx()
	c.Request = req
	c.ResponseWriter = rsp
	controller, err := wkws.HandlerHttpRequest(c)
	if err != nil {
		return
	}
	controller(c)
}

func (wkws *Wkws) HandlerHttpRequest(c *Context) (Controller, error) {
	handler, verify := wkws.VerifyMethod(c.Request.URL.Path, c.Request.Method)
	if !verify {
		ServerFailed(c)
		return nil, &WkwsError{Msg: "error"}
	}
	// TODO Handler Request Params
	return handler, nil
}

func (wkws *Wkws) VerifyMethod(path string, method string) (Controller, bool) {
	for _, r := range wkws.RouterGroup {
		if r.Path == path && r.Method == method {
			return r.Handler, true
		}
	}
	return nil, false
}

func ServerFailed(c *Context) {
	c.ResponseWriter.WriteHeader(405)
	_, err := c.ResponseWriter.Write([]byte("405 method not allowed"))
	if err != nil {
		CLogger("cannot write message %v", err)
	}
	return
}
