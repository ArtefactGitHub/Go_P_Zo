package user

import (
	"net/http"

	util "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/presentation/utils"
	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/user"
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

const resourceKey = "user_id"

func NewFind(uc u.Find) Find {
	return find{find: uc}
}

func (h find) Find(w http.ResponseWriter, r *http.Request, params common.QueryMap) {
	// 指定リソースの取得
	id, err := util.GetResourceId(params, resourceKey)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	user, err := h.find.Do(r.Context(), id)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	myhttp.Write(w, NewResponse(myhttp.NewResponse(nil, http.StatusOK, ""), user), http.StatusOK)
}
