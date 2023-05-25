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
	Delete interface {
		Delete(w http.ResponseWriter, r *http.Request, ps common.QueryMap)
	}
	deleteUsecase struct {
		delete u.Delete
	}
)

func NewDelete(uc u.Delete) Delete {
	return deleteUsecase{delete: uc}
}

func (h deleteUsecase) Delete(w http.ResponseWriter, r *http.Request, params common.QueryMap) {
	// ユーザーIDの取得
	// TODO: userID検証
	_, err := util.GetUserIdFromToken(r.Context())
	if err != nil {
		e := util.Wrap(derr.Unauthorized, fmt.Sprintf("error with GetUserIdFromToken: %#v", err))
		util.HandleError(w, e)
		return
	}

	// 指定リソースの取得
	id, err := util.GetResourceId(params, util.ResourceIdZo)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	err = h.delete.Do(r.Context(), id)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	myhttp.Write(w, NewSimpleResponse(myhttp.NewResponse(nil, http.StatusOK, "")), http.StatusOK)
}
