package client

import (
	"encoding/json"
	"errors"
	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/client"
	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/client"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
	"io"
	"net/http"
)

type (
	Exist interface {
		Post(w http.ResponseWriter, r *http.Request, ps common.QueryMap)
	}
	exist struct {
		exist       u.Exist
		createToken u.CreateToken
	}
	PostRequest struct {
		Id     int    `json:"id"`
		Secret string `json:"secret"`
	}
	PostResponse struct {
		*myhttp.ResponseBase
		AccessToken d.AccessToken `json:"access_token"`
	}
)

func NewPostResponse(res *myhttp.ResponseBase, token d.AccessToken) *PostResponse {
	return &PostResponse{ResponseBase: res, AccessToken: token}
}

func NewExist(e u.Exist, c u.CreateToken) Exist {
	return exist{exist: e, createToken: c}
}

func (h exist) Post(w http.ResponseWriter, r *http.Request, _ common.QueryMap) {
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
		myhttp.Write(w, NewPostResponse(myhttp.NewResponse(errors.New("client not found"), http.StatusNotFound, ""), nil), http.StatusNotFound)
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
