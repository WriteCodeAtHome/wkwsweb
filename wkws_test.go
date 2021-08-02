package wkwsweb

import (
	"log"
	"testing"
)

func PostController(ctx *Context) {
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

func TestWkws_Run(t *testing.T) {
	core := Init()
	core.POST("/post", PostController)
	core.GET("/get", GetController)
	err := core.Run("0.0.0.0", "8081")
	if err != nil {
		log.Println(err)
	}
}

func TestWkws_ContextCache(t *testing.T) {
	core := Init()
	core.GET("/cache", CacheController)
	err := core.Run("0.0.0.0", "8081")
	if err != nil {
		log.Println(err)
	}
}
