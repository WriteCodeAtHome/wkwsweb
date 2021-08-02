package wkwsweb

import "net/http"

type Context struct {
	Request *http.Request
	http.ResponseWriter
	size     int
	status   int
	index    int8
	fullPath string
}
