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

func TestWkws_Run(t *testing.T) {
	core := Init("0.0.0.0", "8081")
	core.POST("/post", PostController)
	core.GET("/get", GetController)
	err := core.Run()
	if err != nil {
		log.Println(err)
	}
}
