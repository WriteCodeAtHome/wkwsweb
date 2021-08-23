package wkwsweb

import "fmt"

const (
	PREFIX = "[WKWS]"
)

func CLogger(s string, v ...interface{}) {
	fmt.Printf(PREFIX+s+"\n", v...)
}
