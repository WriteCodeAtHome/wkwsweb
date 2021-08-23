package wkwsweb

import (
	"net/http"
)

type Error struct {
	Err error
}

type Wkws struct {
	RouterHandlers RouterHandlers
	HttpServer     *http.Server
}

func Init() (core *Wkws) {
	core = &Wkws{
		RouterHandlers: nil,
	}
	return
}

func (wkws *Wkws) Run(add string, port string) (err error) {
	addr := add + ":" + port
	CLogger("HTTP listen the address , %s \n", addr)
	r := http.NewServeMux()
	CLogger("Router map:")
	// Get new routers, filter router path repeat
	registeredRouter := map[string]struct{}{}

	for _, router := range wkws.RouterHandlers {
		if _, exist := registeredRouter[router.Path]; !exist {
			r.HandleFunc(router.Path, wkws.ServeHTTP)
			registeredRouter[router.Path] = struct{}{}
		}
		CLogger("Router Method is %s , Path in %s , Handler is %s", router.Method, router.Path, router.Handler)
	}
	wkws.HttpServer = &http.Server{Addr: addr, Handler: r}
	err = wkws.HttpServer.ListenAndServe()
	return
}

func (wkws *Wkws) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	c := NewCtx()
	c.Request = req
	c.ResponseWriter = rsp
	controller, err := wkws.handlerHttpRequest(c)
	if err != nil {
		return
	}
	controller(c)
}

func (wkws *Wkws) handlerHttpRequest(c *Context) (Controller, error) {
	handler, verify := wkws.handlerVerifyRouter(c.Request.URL.Path, c.Request.Method)
	if !verify {
		ServerMethodsNotAllowed(c)
		return nil, &WkwsError{Msg: "error"}
	}
	// TODO Handler Request Params
	// for that c.Request.URL.Query() , key is value value is key.
	return handler, nil
}

func (wkws *Wkws) handlerVerifyRouter(path string, method string) (Controller, bool) {
	for _, r := range wkws.RouterHandlers {
		if r.Path == path && r.Method == method {
			return r.Handler, true
		}
	}
	return nil, false
}

func ServerMethodsNotAllowed(c *Context) {
	c.ResponseWriter.WriteHeader(405)
	_, err := c.ResponseWriter.Write([]byte("405 method not allowed"))
	if err != nil {
		CLogger("cannot write message %v", err)
	}
	return
}
