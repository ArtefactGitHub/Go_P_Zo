package category

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	derr "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
	util "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/presentation/utils"
	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/zo/category"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

type (
	Create interface {
		Create(w http.ResponseWriter, r *http.Request, ps common.QueryMap)
	}
	create struct {
		create u.Create
	}
)

func NewCreate(uc u.Create) Create {
	return create{create: uc}
}

func (h create) Create(w http.ResponseWriter, r *http.Request, _ common.QueryMap) {
	// ユーザーIDの取得
	userID, err := util.GetUserIdFromToken(r.Context())
	if err != nil {
		e := util.Wrap(derr.Unauthorized, fmt.Sprintf("error with GetUserIdFromToken: %#v", err))
		util.HandleError(w, e)
		return
	}

	// リクエスト情報からモデルを生成
	m, err := contentToCreateModel(r, userID)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	v, err := h.create.Do(r.Context(), m)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	myhttp.Write(w, NewResponse(myhttp.NewResponse(nil, http.StatusOK, ""), v), http.StatusOK)
}

func contentToCreateModel(r *http.Request, userID int) (d.Category, error) {
	body := make([]byte, r.ContentLength)
	if _, err := r.Body.Read(body); err != nil && err != io.EOF {
		return nil, err
	}
	var req PostRequest
	err := json.Unmarshal(body, &req)
	if err != nil {
		return nil, err
	}

	return d.NewCategory(0, req.Name, userID), nil
}
