package zo

import (
	"fmt"
	"net/http"

	derr "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
	util "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/presentation/utils"
	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/zo"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

type (
	Find interface {
		Find(w http.ResponseWriter, r *http.Request, ps common.QueryMap)
	}
	find struct {
		find u.Find
	}
)

const resourceKey = "zo_id"

func NewFind(uc u.Find) Find {
	return find{find: uc}
}

func (h find) Find(w http.ResponseWriter, r *http.Request, params common.QueryMap) {
	// ユーザーIDの取得
	userId, err := util.GetUserIdFromToken(r.Context())
	if err != nil {
		e := util.Wrap(derr.Unauthorized, fmt.Sprintf("error with GetUserIdFromToken: %#v", err))
		util.HandleError(w, e)
		return
	}

	// 指定リソースの取得
	id, err := util.GetResourceId(params, resourceKey)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	zo, err := h.find.Do(r.Context(), id)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	// 非リソース所有者の場合
	if userId != zo.UserId {
		e := util.Wrap(derr.Unauthorized, fmt.Sprintf("difference userID. request: %d, resource: %d", userId, zo.UserId))
		util.HandleError(w, e)
		return
	}

	myhttp.Write(w, NewZoResponse(myhttp.NewResponse(nil, http.StatusOK, ""), zo), http.StatusOK)
}
