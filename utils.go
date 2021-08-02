package wkwsweb

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

type WkwsError struct {
	Msg string
}

func (e *WkwsError) Error() string {
	return e.Msg
}
