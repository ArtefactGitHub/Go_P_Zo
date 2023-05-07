package error

import (
	"errors"
)

var (
	NotFound   = errors.New("not found")
	BadRequest = errors.New("bad request")
)
