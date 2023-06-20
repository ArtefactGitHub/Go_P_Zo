package session

import (
	"net/http"

	i "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/session"
	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/session"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

var (
	us = u.NewLogin(i.NewRepository())
	h  = NewLogin(us)
)

var Routes = map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap){
	{Path: "/api/v2/login", Method: "POST", NeedAuth: true}: h.Login,
}
