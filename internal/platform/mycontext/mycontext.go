package mycontext

import (
	"context"
	"fmt"
	"runtime"
)

type ContextKey string

const AuthorizedKey ContextKey = "authorized"
const UserTokenKey ContextKey = "user_token"

func NewContext(parent context.Context, key interface{}, value interface{}) context.Context {
	return context.WithValue(parent, key, value)
}

func FromContextBool(ctx context.Context, key interface{}) (bool, error) {
	v := ctx.Value(key)
	value, ok := v.(bool)
	if !ok {
		return false, fmt.Errorf("%s header not found", key)
	}

	return value, nil
}

func FromContextStr(ctx context.Context, key interface{}) (string, error) {
	v := ctx.Value(key)
	value, ok := v.(string)
	if !ok {
		fmt.Println(runtime.Caller(1))
		return "", fmt.Errorf("%s header not found", key)
	}

	return value, nil
}
