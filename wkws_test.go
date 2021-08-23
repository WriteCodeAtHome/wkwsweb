package wkwsweb

import (
	"log"
	"net/http/httptest"
	"testing"
)

func PostController(ctx *Context) {
	log.Println(ctx.GetMap())
	log.Println(ctx.Request.Method)
	log.Println(ctx.Request.Header["Content-Type"])
	log.Println(ctx.GetRequestContext())
	ctx.ResponseWriter.Write([]byte("This is Post Controller"))
}

func GetController(ctx *Context) {
	ctx.ResponseWriter.Write([]byte("This is Get Controller"))
}

func CacheController(ctx *Context) {
	ctx.Set("Hello", "World")
	bytes := []byte("Cache : " + ctx.Get("Hello").(string))
	ctx.ResponseWriter.Write(bytes)
}

func ResponseController(ctx *Context) {
	ctx.ResponseJSON(H{
		"HelloWorld": "Millyn",
	})
}

func TestWkws_HttpRun(t *testing.T) {
	core := Init()
	core.GET("/cache", CacheController)
	core.GET("/response", ResponseController)
	core.POST("/response", ResponseController)

	var ts *httptest.Server

	ts = httptest.NewServer(core)

	defer ts.Close()
}
