package error

import (
	"errors"
	"net/http"
)

var (
	BadRequest   = errors.New(http.StatusText(http.StatusBadRequest))
	Unauthorized = errors.New(http.StatusText(http.StatusUnauthorized))
	NotFound     = errors.New(http.StatusText(http.StatusNotFound))
	Conflict     = errors.New(http.StatusText(http.StatusConflict))
)
