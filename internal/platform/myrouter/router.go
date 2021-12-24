package myrouter

import (
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/zo"
)

type MyRouter struct {
}

func (r *MyRouter) Routing(mux *http.ServeMux) {
	zo.Routing(mux)
}
