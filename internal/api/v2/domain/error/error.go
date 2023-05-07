package error

import (
	"errors"
)

var (
	BadRequest   = errors.New("bad request")
	Unauthorized = errors.New("unauthorized")
	NotFound     = errors.New("not found")
)
