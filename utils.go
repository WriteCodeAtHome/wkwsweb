package wkwsweb

type WkwsError struct {
	Msg string
}

func (e *WkwsError) Error() string {
	return e.Msg
}
