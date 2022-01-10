package auth

import (
	"net/http"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
	"github.com/golang-jwt/jwt"
)

type authController struct {
}

func (c *authController) post(w http.ResponseWriter, r *http.Request, ps common.QueryMap) {
	claims := myauth.AuthClaims{StandardClaims: &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(config.Cfg.Auth.TokenExpiration)).Unix(),
		Issuer:    "zo.auth.service",
	},
		TokenType: "access",
	}

	// https://pkg.go.dev/github.com/golang-jwt/jwt#NewWithClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := token.SignedString([]byte(config.Cfg.Auth.SignKey))
	if err != nil {
		myhttp.WriteError(w, err, http.StatusBadRequest, "")
	}

	res := PostResponse{
		ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusOK, Error: nil},
		AccessToken:  AccessToken{Jwt: jwt, ExpiresAt: claims.ExpiresAt}}
	myhttp.WriteSuccess(w, res, http.StatusOK)
}
