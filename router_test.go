package wkwsweb

import "testing"

func TestRouters(t *testing.T) {
	wkws := Init("0.0.0.0", "8081")
	wkws.POST("/hello", wkws.ServeHTTP)
	err := wkws.Run()
	if err != nil {
		return
	}
}
