package zo

import (
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
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mytime"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

type (
	Update interface {
		Update(w http.ResponseWriter, r *http.Request, ps common.QueryMap)
	}
	update struct {
		update u.Update
	}
)

func NewUpdate(uc u.Update) Update {
	return update{update: uc}
}

func (h update) Update(w http.ResponseWriter, r *http.Request, params common.QueryMap) {
	// ユーザーIDの取得
	userID, err := util.GetUserIdFromToken(r.Context())
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

	// リクエスト情報からモデルを生成
	m, err := contentToUpdateModel(r, id, userID)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	z, err := h.update.Do(r.Context(), m)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	myhttp.Write(w, NewGetResponse(myhttp.NewResponse(nil, http.StatusOK, ""), z), http.StatusOK)
}

func contentToUpdateModel(r *http.Request, id int, userID int) (zo.Zo, error) {
	body := make([]byte, r.ContentLength)
	if _, err := r.Body.Read(body); err != nil && err != io.EOF {
		return zo.Zo{}, err
	}
	var req PostRequest
	err := json.Unmarshal(body, &req)
	if err != nil {
		return zo.Zo{}, err
	}

	return zo.NewZo(
			id,
			req.AchievementDate,
			req.Exp,
			req.CategoryId,
			req.Message,
			time.Now(),
			mytime.NullTime(time.Now()),
			userID),
		nil
}
