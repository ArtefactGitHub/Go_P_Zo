package zo

import (
	"net/http"

	i "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/zo"
	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/zo"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

var (
	uf = u.NewFind(i.NewRepository())
	uc = u.NewCreate(i.NewRepository())
	uu = u.NewUpdate(i.NewRepository())
	hf = NewFind(uf)
	hc = NewCreate(uc)
	hu = NewUpdate(uu)
)

var Routes = map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap){
	{Path: "/api/v2/zos/:zo_id", Method: "GET", NeedAuth: true}: hf.Find,
	{Path: "/api/v2/zos", Method: "POST", NeedAuth: true}:       hc.Create,
	{Path: "/api/v2/zos/:zo_id", Method: "PUT", NeedAuth: true}: hu.Update,
}
