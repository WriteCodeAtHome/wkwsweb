package wkwsweb

import (
	"net/http"
)

type Error struct {
	Err error
}

type Wkws struct {
	Address     string `json:"address"`
	Port        string `json:"port"`
	RouterGroup Routers
}

func Init(add string, port string) (core *Wkws) {
	core = &Wkws{
		Address:     add,
		Port:        port,
		RouterGroup: nil,
	}
	return
}

func (wkws *Wkws) Run() (err error) {
	addr := wkws.Address + ":" + wkws.Port
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
	c := Context{}
	c.Request = req
	c.ResponseWriter = rsp
	controller, err := wkws.CheckMethod(&c)
	if err != nil {
		return
	}
	controller(&c)
}

func (wkws *Wkws) CheckMethod(c *Context) (Controller, error) {
	handler, verify := wkws.VerifyMethod(c.Request.RequestURI, c.Request.Method)
	if !verify {
		ServerFailed(c)
		return nil, &WkwsError{Msg: "error"}
	}
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
