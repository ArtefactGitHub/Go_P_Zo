package mycontext

import (
	"context"
	"fmt"
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
		return false, fmt.Errorf("%s not found", key)
	}

	return value, nil
}
