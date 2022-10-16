package client

import (
	"net/http"

	i "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/client"
	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/client"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

var (
	ue = u.NewExist(i.NewRepository())
	uc = u.NewCreateToken()
	h  = NewCreateToken(ue, uc)
)

var Routes = map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap){
	{Path: "/api/v2/client/token", Method: "POST", NeedAuth: false}: h.Create,
}
