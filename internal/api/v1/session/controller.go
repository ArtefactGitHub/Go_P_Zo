package session

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

type sessionController struct {
	s SessionService
}

func (c *sessionController) post(w http.ResponseWriter, r *http.Request, ps common.QueryMap) {
	// リクエスト情報からモデルを生成
	m, err := c.contentToModel(r)
	log.Printf("contentToModel: %v", m)
	if err != nil {
		myhttp.Write(w, NewPostResponse(myhttp.NewResponse(err, http.StatusInternalServerError, ""), nil), http.StatusInternalServerError)
		return
	}

	sessionData, err := c.s.Login(r.Context(), m.Identifier, m.Secret)
	if err != nil {
		myhttp.Write(w, NewPostResponse(myhttp.NewResponse(err, http.StatusInternalServerError, ""), nil), http.StatusInternalServerError)
		return
	}
	if sessionData == nil {
		myhttp.Write(w, NewPostResponse(myhttp.NewResponse(errors.New("not found"), http.StatusNotFound, ""), nil), http.StatusNotFound)
		return
	}

	myhttp.Write(w, NewPostResponse(myhttp.NewResponse(nil, http.StatusOK, ""), sessionData), http.StatusOK)
}

// リクエスト情報からモデルの生成
func (c *sessionController) contentToModel(r *http.Request) (*SessionRequest, error) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var result SessionRequest
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
