package wkwsweb

import "net/http"

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
