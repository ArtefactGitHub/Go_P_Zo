package middleware

import (
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/auth"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/user"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/zo"
)

type RouterMiddleware struct {
	r    http.Handler
	next IMiddleware
}

func NewRouterMiddleware() IMiddleware {
	return &RouterMiddleware{r: myrouter.NewMyRouterWithRoutes(
		auth.Routes,
		zo.Routes,
		user.Routes)}
}

func (m *RouterMiddleware) SetNext(next IMiddleware) {
	m.next = next
}

func (m *RouterMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	m.r.ServeHTTP(w, req)
	if m.next != nil {
		m.next.ServeHTTP(w, req)
	}
}
