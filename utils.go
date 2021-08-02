package wkwsweb

func ServerFailed(c *Context) {
	c.ResponseWriter.WriteHeader(405)
	_, err := c.ResponseWriter.Write([]byte("405 method not allowed"))
	if err != nil {
		CLogger("cannot write message %v", err)
	}
	return
}
