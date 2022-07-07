package utils

type Error struct {
	Msg   string `json:"msg"`
	Error error  `json:"-"`
}

func NewError(msg string, error error) *Error {
	return &Error{
		Msg:   msg,
		Error: error,
	}
}
