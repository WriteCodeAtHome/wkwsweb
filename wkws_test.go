package wkwsweb

import (
	"log"
	"net/http"
	"testing"
)

func PostController(rsp http.ResponseWriter, req *http.Request) {
	_, err := rsp.Write([]byte("The is POST method"))
	if err != nil {
		return
	}
}

func GetController(rsp http.ResponseWriter, req *http.Request) {
	_, err := rsp.Write([]byte("The is GET method"))
	if err != nil {
		return
	}
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
