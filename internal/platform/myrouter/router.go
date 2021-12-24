package myrouter

import (
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/zo"
	"github.com/julienschmidt/httprouter"
)

type Router struct {
	r httprouter.Router
}

func New() *Router {
	return &Router{r: *httprouter.New()}
}

func (mr *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	mr.r.ServeHTTP(w, req)
}

func (mr *Router) Routing() {
	zo.Routing(&mr.r)
}
