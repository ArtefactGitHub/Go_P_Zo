package zo

import (
	"net/http"

	i "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/zo"
	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/zo"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

var (
	uc = u.NewFind(i.NewRepository())
	h  = NewFind(uc)
)

var Routes = map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap){
	{Path: "/api/v2/zos/:zo_id", Method: "GET", NeedAuth: true}: h.Find,
}
