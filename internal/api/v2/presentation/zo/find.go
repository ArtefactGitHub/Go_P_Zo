package zo

import (
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

func NewFind(uc u.Find) Find {
	return find{find: uc}
}

func (h find) Find(w http.ResponseWriter, r *http.Request, params common.QueryMap) {
	// ユーザーIDの取得
	userId, err := util.GetUserIdFromToken(r.Context())
	if err != nil {
		util.HandleError(w, derr.Unauthorized)
		return
	}

	// 指定リソースの取得
	id, err := util.GetResourceId(params, util.ResourceIdZo)
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
		util.HandleError(w, derr.Unauthorized)
		return
	}

	myhttp.Write(w, NewGetResponse(myhttp.NewResponse(nil, http.StatusOK, ""), zo), http.StatusOK)
}
