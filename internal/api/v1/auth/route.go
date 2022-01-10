package auth

import (
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

var uc authController = authController{}

var Routes map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap) = map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap){
	{Path: "/api/v1/auth/token", Method: "POST", NeedAuth: false}: uc.post,
}
