package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mycontext"
)

type JwtMiddleware struct {
	next IMiddleware
}

func NewJwtMiddleware() IMiddleware {
	return &JwtMiddleware{}
}

func (m *JwtMiddleware) SetNext(next IMiddleware) {
	m.next = next
}

func (m *JwtMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	token := req.Header.Get("Authorization")
	extractedToken := strings.Split(token, "Bearer ")

	ctx := req.Context()
	if len(extractedToken) == 2 {
		token = strings.TrimSpace(extractedToken[1])
		log.Printf("token: %v", token)

		// TODO: validate

		// set value to context
		ctx = mycontext.NewContext(ctx, mycontext.AuthorizedKey, true)
	}

	if m.next != nil {
		m.next.ServeHTTP(w, req.WithContext(ctx))
	}
}
