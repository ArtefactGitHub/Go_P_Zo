package myrouter

import (
	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/zo"
)

type MyRouter struct {
}

func (r *MyRouter) Routing() {
	zo.Routing()
}
