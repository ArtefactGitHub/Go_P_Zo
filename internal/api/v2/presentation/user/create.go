package user

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"time"

	domain "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/user"
	util "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/presentation/utils"
	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/user"
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

	PostRequest struct {
		GivenName  string `json:"given_name"`
		FamilyName string `json:"family_name"`
		Email      string `json:"email"`
		Password   string `json:"password"`
	}
)

func NewCreate(uc u.Create) Create {
	return create{create: uc}
}

func (h create) Create(w http.ResponseWriter, r *http.Request, _ common.QueryMap) {
	// リクエスト情報からモデルを生成
	v, err := contentToCreateModel(r)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	user, err := h.create.Do(r.Context(), v)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	myhttp.Write(w, NewResponse(myhttp.NewResponse(nil, http.StatusCreated, ""), user), http.StatusOK)
}

// リクエスト情報からモデルの生成
func contentToCreateModel(r *http.Request) (domain.User, error) {
	body := make([]byte, r.ContentLength)
	if _, err := r.Body.Read(body); err != nil && err != io.EOF {
		return domain.User{}, err
	}
	var result PostRequest
	err := json.Unmarshal(body, &result)
	if err != nil {
		return domain.User{}, err
	}

	return domain.NewUser(
			0,
			result.GivenName,
			result.FamilyName,
			result.Email,
			result.Password,
			time.Time{},
			sql.NullTime{}),
		nil
}
