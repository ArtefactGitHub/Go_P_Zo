package client

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/client"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

type (
	CreateToken interface {
		Create(w http.ResponseWriter, r *http.Request, ps common.QueryMap)
	}
	createToken struct {
		exist       u.Exist
		createToken u.CreateToken
	}
)

func NewCreateToken(e u.Exist, c u.CreateToken) CreateToken {
	return createToken{exist: e, createToken: c}
}

func (h createToken) Create(w http.ResponseWriter, r *http.Request, _ common.QueryMap) {
	// リクエスト情報からモデルを生成
	m, err := contentToModel(r)
	if err != nil {
		myhttp.Write(w, NewPostResponse(myhttp.NewResponse(err, http.StatusInternalServerError, ""), nil), http.StatusInternalServerError)
		return
	}

	exist, err := h.exist.Do(r.Context(), m.Id, m.Secret)
	if err != nil {
		myhttp.Write(w, NewPostResponse(myhttp.NewResponse(err, http.StatusInternalServerError, ""), nil), http.StatusInternalServerError)
		return
	}
	if !exist {
		myhttp.Write(w, NewPostResponse(myhttp.NewResponse(errors.New("invalid identifier"), http.StatusUnauthorized, ""), nil), http.StatusUnauthorized)
		return
	}

	token, err := h.createToken.Do()
	if err != nil {
		myhttp.Write(w, NewPostResponse(myhttp.NewResponse(err, http.StatusInternalServerError, ""), nil), http.StatusInternalServerError)
		return
	}

	myhttp.Write(w, NewPostResponse(myhttp.NewResponse(nil, http.StatusOK, ""), token), http.StatusOK)
}

// リクエスト情報からモデルの生成
func contentToModel(r *http.Request) (PostRequest, error) {
	body := make([]byte, r.ContentLength)
	if _, err := r.Body.Read(body); err != nil && err != io.EOF {
		return PostRequest{}, err
	}
	var result PostRequest
	err := json.Unmarshal(body, &result)
	if err != nil {
		return PostRequest{}, err
	}
	return result, nil
}
