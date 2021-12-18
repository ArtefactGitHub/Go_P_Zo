package myerror

import "fmt"

type myError struct {
	Message string `json:"message"`
	Origin  error  `json:"-"`
}

func (e *myError) Error() string {
	return e.Message
}

func NewError(err error, desc string) error {
	switch {
	case err == nil:
		return &myError{Message: desc, Origin: nil}
	case desc == "":
		return &myError{Message: err.Error(), Origin: err}
	default:
		return &myError{Message: fmt.Sprintf("%s: %s", desc, err.Error()), Origin: err}
	}
}
