package client

import (
	domain "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/client"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
	"github.com/golang-jwt/jwt"
)

type (
	CreateToken interface {
		Do() (domain.AccessToken, error)
	}
	createToken struct{}
)

func NewCreateToken() CreateToken {
	return createToken{}
}

func (u createToken) Do() (domain.AccessToken, error) {
	return u.create()
}

func (u createToken) create() (domain.AccessToken, error) {
	claims := myauth.AuthClaims{StandardClaims: &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(config.Cfg.Auth.TokenExpiration)).Unix(),
		Issuer:    domain.Issuer,
	},
		TokenType: domain.TokenType,
	}

	// https://pkg.go.dev/github.com/golang-jwt/jwt#NewWithClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	j, err := token.SignedString([]byte(config.Cfg.Auth.SignKey))
	if err != nil {
		return nil, err
	}

	return domain.NewAccessToken(j, claims.ExpiresAt), nil
}
