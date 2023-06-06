package auth

import (
	"net/http"

	i "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/auth"
	iu "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/user"
	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/auth"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

var (
	uc = u.NewCreate(i.NewRepository(), iu.NewRepository())
	h  = NewCreateToken(uc)
)

var Routes = map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap){
	{Path: "/api/v2/usertokens", Method: http.MethodPost, NeedAuth: false}: h.Create,
}
