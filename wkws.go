package wkwsweb

import (
	"net/http"
)

type Wkws struct {
	Address string `json:"address"`
	Port    string `json:"port"`
	Router  *http.ServeMux
}

func Init(add string, port string) (core *Wkws) {
	core = &Wkws{
		Address: add,
		Port:    port,
		Router:  http.NewServeMux(),
	}
	return
}

func (wkws *Wkws) Run() (err error) {
	addr := wkws.Address + ":" + wkws.Port
	CLogger("HTTP listen the address , %s \n", addr)
	err = http.ListenAndServe(addr, wkws.Router)
	return
}
