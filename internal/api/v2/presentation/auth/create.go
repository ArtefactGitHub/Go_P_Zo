package auth

import (
	"encoding/json"
	"io"
	"net/http"

	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/auth"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

type (
	CreateToken interface {
		Create(w http.ResponseWriter, r *http.Request, ps common.QueryMap)
	}
	createToken struct {
		createToken u.Create
	}
)

func NewCreateToken(c u.Create) CreateToken {
	return createToken{createToken: c}
}

func (h createToken) Create(w http.ResponseWriter, r *http.Request, _ common.QueryMap) {
	// リクエスト情報からモデルを生成
	m, err := contentToModel(r)
	if err != nil {
		myhttp.Write(w, NewPostResponse(myhttp.NewResponse(err, http.StatusInternalServerError, ""), nil), http.StatusInternalServerError)
		return
	}

	token, err := h.createToken.Do(r.Context(), m)
	if err != nil {
		myhttp.Write(w, NewPostResponse(myhttp.NewResponse(err, http.StatusInternalServerError, ""), nil), http.StatusInternalServerError)
		return
	}

	myhttp.Write(w, NewPostResponse(myhttp.NewResponse(nil, http.StatusOK, ""), token), http.StatusOK)
}

// リクエスト情報からモデルの生成
func contentToModel(r *http.Request) (u.CreateTokenData, error) {
	body := make([]byte, r.ContentLength)
	if _, err := r.Body.Read(body); err != nil && err != io.EOF {
		return u.CreateTokenData{}, err
	}
	var result u.CreateTokenData
	err := json.Unmarshal(body, &result)
	if err != nil {
		return u.CreateTokenData{}, err
	}
	return result, nil
}
