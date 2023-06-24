package user

import (
	"net/http"

	infra "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/user"
	usecase "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/user"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

var (
	uf = usecase.NewFind(infra.NewRepository())
	uc = usecase.NewCreate(infra.NewRepository())
	hf = NewFind(uf)
	hc = NewCreate(uc)
)

var Routes = map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap){
	{Path: "/api/v2/users/:user_id", Method: http.MethodGet, NeedAuth: false}: hf.Find,
	{Path: "/api/v2/users", Method: http.MethodPost, NeedAuth: false}:         hc.Create,
}
