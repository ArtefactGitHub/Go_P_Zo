package zo

import (
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

var zc zoController = zoController{}

var Routes map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap) = map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap){
	{Path: "/api/v1/zos", Method: "GET", NeedAuth: true}:           zc.getAll,
	{Path: "/api/v1/zos/:zo_id", Method: "GET", NeedAuth: true}:    zc.get,
	{Path: "/api/v1/zos", Method: "POST", NeedAuth: true}:          zc.post,
	{Path: "/api/v1/zos/:zo_id", Method: "PUT", NeedAuth: true}:    zc.update,
	{Path: "/api/v1/zos/:zo_id", Method: "DELETE", NeedAuth: true}: zc.delete,
}
