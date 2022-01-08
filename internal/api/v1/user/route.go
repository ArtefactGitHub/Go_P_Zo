package user

import (
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

var uc userController = userController{}

var Routes map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap) = map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap){
	{Path: "/api/v1/users", Method: "GET"}:             uc.getAll,
	{Path: "/api/v1/users/:user_id", Method: "GET"}:    uc.get,
	{Path: "/api/v1/users", Method: "POST"}:            uc.post,
	{Path: "/api/v1/users/:user_id", Method: "PUT"}:    uc.update,
	{Path: "/api/v1/users/:user_id", Method: "DELETE"}: uc.delete,
}
