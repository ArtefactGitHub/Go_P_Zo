package middleware

import (
	"log"
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mycontext"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
)

var headerNames map[string]mycontext.ContextKey = map[string]mycontext.ContextKey{
	myhttp.UserTokenName: mycontext.UserTokenKey,
}

type HeaderMiddleware struct {
	next IMiddleware
}

func NewHeaderMiddleware() (IMiddleware, error) {
	return &HeaderMiddleware{}, nil
}

func (m *HeaderMiddleware) SetNext(next IMiddleware) {
	m.next = next
}

func (m *HeaderMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	for name, contextKey := range headerNames {
		val := req.Header.Get(name)

		if val != "" {
			log.Printf("[HeaderMiddleware] %s: %s", name, val)
			ctx = mycontext.NewContext(ctx, contextKey, val)
		}
	}

	if m.next != nil {
		m.next.ServeHTTP(w, req.WithContext(ctx))
	}
}
