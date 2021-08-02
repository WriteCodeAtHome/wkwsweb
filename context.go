package wkwsweb

import (
	"encoding/json"
	"net/http"
)

type Param struct {
	Key   string
	Value string
}

type Params []Param

type Context struct {
	Request *http.Request
	http.ResponseWriter
	Cache  map[string]interface{}
	Params Params
}

type Controller func(ctx *Context)

func NewCtx() *Context {
	return &Context{
		nil,
		nil,
		map[string]interface{}{},
		make(Params, 0, 100),
	}
}

func (ps Params) Get(name string) (string, bool) {
	for _, entry := range ps {
		if entry.Key == name {
			return entry.Value, true
		}
	}
	return "", false
}

func (ctx *Context) Set(key string, value interface{}) {
	ctx.Cache[key] = value
}

func (ctx *Context) Get(key string) interface{} {
	value := ctx.Cache[key]
	return value
}

func (ctx *Context) ResponseJSON(data interface{}) {
	ctx.ResponseWriter.Header().Add("Content-type", "application/json")
	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}
	_, err = ctx.ResponseWriter.Write(marshal)
	if err != nil {
		CLogger("Cannot write response json %v", err)
		return
	}
}
