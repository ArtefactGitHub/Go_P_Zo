package zo

import (
	"net/http"

	i "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/zo"
	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/zo"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

var (
	ufs = u.NewFinds(i.NewRepository())
	uf  = u.NewFind(i.NewRepository())
	uc  = u.NewCreate(i.NewRepository())
	uu  = u.NewUpdate(i.NewRepository())
	ud  = u.NewDelete(i.NewRepository())
	hfs = NewFinds(ufs)
	hf  = NewFind(uf)
	hc  = NewCreate(uc)
	hu  = NewUpdate(uu)
	hd  = NewDelete(ud)
)

var Routes = map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap){
	{Path: "/api/v2/zos", Method: http.MethodGet, NeedAuth: true}:           hfs.Finds,
	{Path: "/api/v2/zos/:zo_id", Method: http.MethodGet, NeedAuth: true}:    hf.Find,
	{Path: "/api/v2/zos", Method: http.MethodPost, NeedAuth: true}:          hc.Create,
	{Path: "/api/v2/zos/:zo_id", Method: http.MethodPut, NeedAuth: true}:    hu.Update,
	{Path: "/api/v2/zos/:zo_id", Method: http.MethodDelete, NeedAuth: true}: hd.Delete,
}
