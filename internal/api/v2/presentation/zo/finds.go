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
	Finds interface {
		Finds(w http.ResponseWriter, r *http.Request, ps common.QueryMap)
	}
	finds struct {
		finds u.Finds
	}
)

func NewFinds(uc u.Finds) Finds {
	return finds{finds: uc}
}

func (h finds) Finds(w http.ResponseWriter, r *http.Request, _ common.QueryMap) {
	// ユーザーIDの取得
	userId, err := util.GetUserIdFromToken(r.Context())
	if err != nil {
		e := util.Wrap(derr.Unauthorized, fmt.Sprintf("error with GetUserIdFromToken: %#v", err))
		util.HandleError(w, e)
		return
	}

	result, err := h.finds.Do(r.Context(), userId)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	myhttp.Write(w, NewZosResponse(myhttp.NewResponse(nil, http.StatusOK, ""), result), http.StatusOK)
}
