package session

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	e "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/session"
	u "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/session"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

type (
	Login interface {
		Login(w http.ResponseWriter, r *http.Request, ps common.QueryMap)
	}
	login struct {
		login u.Login
	}
)

func NewLogin(u u.Login) Login {
	return login{login: u}
}

func (h login) Login(w http.ResponseWriter, r *http.Request, _ common.QueryMap) {
	// リクエスト情報からモデルを生成
	m, err := contentToModel(r)
	if err != nil {
		statusCode := ToStatusCode(err)
		myhttp.Write(w, NewPostResponse(myhttp.NewResponse(err, statusCode, ""), d.SessionData{}), statusCode)
		return
	}

	data, err := h.login.Do(r.Context(), m.Identifier, m.Secret)
	if err != nil {
		statusCode := ToStatusCode(err)
		myhttp.Write(w, NewPostResponse(myhttp.NewResponse(err, statusCode, ""), d.SessionData{}), statusCode)
		return
	}

	myhttp.Write(w, NewPostResponse(myhttp.NewResponse(nil, http.StatusOK, ""), data), http.StatusOK)
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

func ToStatusCode(err error) int {
	switch {
	case errors.Is(err, e.NotFound):
		return http.StatusNotFound
	case errors.Is(err, e.BadRequest):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
