package user

import (
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

var uc userController = userController{}
var utc userTokenController = userTokenController{}
var ucc userCategoryController = userCategoryController{}

var Routes map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap) = map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap){
	{Path: "/api/v1/users", Method: "GET", NeedAuth: true}:                     uc.getAll,
	{Path: "/api/v1/users/:user_id", Method: "GET", NeedAuth: true}:            uc.get,
	{Path: "/api/v1/users", Method: "POST", NeedAuth: true}:                    uc.post,
	{Path: "/api/v1/users/:user_id", Method: "PUT", NeedAuth: true}:            uc.update,
	{Path: "/api/v1/users/:user_id", Method: "DELETE", NeedAuth: true}:         uc.delete,
	{Path: "/api/v1/usertokens", Method: "POST", NeedAuth: false}:              utc.post,
	{Path: "/api/v1/users/:user_id/categories", Method: "GET", NeedAuth: true}: ucc.getAll,
}
