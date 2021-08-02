package wkwsweb

import "net/http"

type Context struct {
	Request *http.Request
	http.ResponseWriter
}

type Controller func(ctx *Context)
