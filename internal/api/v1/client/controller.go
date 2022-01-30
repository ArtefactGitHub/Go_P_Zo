package client

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
	"github.com/golang-jwt/jwt"
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

	token, err := createAccessToken()
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

func createAccessToken() (*AccessToken, error) {
	claims := myauth.AuthClaims{StandardClaims: &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(config.Cfg.Auth.TokenExpiration)).Unix(),
		Issuer:    "zo.auth.service",
	},
		TokenType: "access",
	}

	// https://pkg.go.dev/github.com/golang-jwt/jwt#NewWithClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := token.SignedString([]byte(config.Cfg.Auth.SignKey))
	log.Printf("signed: %v", jwt)

	if err != nil {
		return nil, err
	}

	return &AccessToken{Jwt: jwt, ExpiresAt: claims.ExpiresAt}, nil
}
