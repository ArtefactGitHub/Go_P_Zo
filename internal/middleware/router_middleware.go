package middleware

import (
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

type RouterMiddleware struct {
	r    http.Handler
	next IMiddleware
}

func NewRouterMiddleware(routes ...map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap)) IMiddleware {
	return &RouterMiddleware{r: myrouter.NewMyRouterWithRoutes(
		routes...)}
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
