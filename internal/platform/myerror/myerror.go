package myerror

import "fmt"

type Error struct {
	Message string `json:"message"`
	Origin  error  `json:"-"`
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(err error, desc string) *Error {
	switch {
	case err == nil:
		return nil
	case desc == "":
		return &Error{Message: err.Error(), Origin: err}
	default:
		return &Error{Message: fmt.Sprintf("%s: %s", desc, err.Error()), Origin: err}
	}
}
