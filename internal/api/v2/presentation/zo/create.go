package zo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	derr "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
	util "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/presentation/utils"
	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/zo"
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

	z, err := h.create.Do(r.Context(), m)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	myhttp.Write(w, NewZoResponse(myhttp.NewResponse(nil, http.StatusOK, ""), z), http.StatusOK)
}

func contentToCreateModel(r *http.Request, userID int) (zo.Zo, error) {
	body := make([]byte, r.ContentLength)
	if _, err := r.Body.Read(body); err != nil && err != io.EOF {
		return zo.Zo{}, err
	}
	var req PostRequest
	err := json.Unmarshal(body, &req)
	if err != nil {
		return zo.Zo{}, err
	}

	return zo.NewZo(0, req.AchievementDate, req.Exp, req.CategoryId, req.Message, time.Now(), sql.NullTime{}, userID), nil
}
