package middleware

import (
	"net/http"
)

type IMiddleware interface {
	SetNext(IMiddleware)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

func CreateHandler(mids ...IMiddleware) http.Handler {
	len := len(mids)
	for i, m := range mids {
		if i+1 < len {
			m.SetNext(mids[i+1])
		}
	}
	return mids[0]
}
