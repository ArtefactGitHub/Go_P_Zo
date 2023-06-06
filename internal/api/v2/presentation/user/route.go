package user

import (
	"net/http"

	i "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/user"
	uc "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/user"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

var (
	uf = uc.NewFind(i.NewRepository())
	hf = NewFind(uf)
)

var Routes = map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap){
	{Path: "/api/v2/users/:user_id", Method: http.MethodGet, NeedAuth: false}: hf.Find,
}
