package mycontext

import (
	"context"
	"errors"
)

type contextKey string

const AuthorizedKey contextKey = "authorized"

func NewContext(parent context.Context, key interface{}, value interface{}) context.Context {
	return context.WithValue(parent, key, value)
}

func FromContextBool(ctx context.Context, key interface{}) (bool, error) {
	v := ctx.Value(key)
	value, ok := v.(bool)
	if !ok {
		return false, errors.New("value not found")
	}

	return value, nil
}
