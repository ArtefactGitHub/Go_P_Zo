package client

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

type clientController struct {
	s clientService
}

func (c *clientController) post(w http.ResponseWriter, r *http.Request, ps common.QueryMap) {
	// リクエスト情報からモデルを生成
	m, err := c.contentToModel(r)
	log.Printf("contentToModel: %v", m)
	if err != nil {
		myhttp.Write(w, NewPostResponse(myhttp.NewResponse(err, http.StatusInternalServerError, ""), nil), http.StatusInternalServerError)
		return
	}

	exist, err := c.s.Exist(r.Context(), m.Id, m.Secret)
	if err != nil {
		myhttp.Write(w, NewPostResponse(myhttp.NewResponse(err, http.StatusInternalServerError, ""), nil), http.StatusInternalServerError)
		return
	}
	if !exist {
		myhttp.Write(w, NewPostResponse(myhttp.NewResponse(errors.New("client not found"), http.StatusNotFound, ""), nil), http.StatusNotFound)
		return
	}

	token, err := c.s.createAccessToken()
	if err != nil {
		myhttp.Write(w, NewPostResponse(myhttp.NewResponse(err, http.StatusInternalServerError, ""), nil), http.StatusInternalServerError)
		return
	}

	myhttp.Write(w, NewPostResponse(myhttp.NewResponse(nil, http.StatusOK, ""), token), http.StatusOK)
}

// リクエスト情報からモデルの生成
func (c *clientController) contentToModel(r *http.Request) (*Client, error) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var result Client
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
