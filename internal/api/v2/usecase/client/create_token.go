package client

import (
	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/client"
	"log"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
	"github.com/golang-jwt/jwt"
)

type (
	CreateToken interface {
		Do() (d.AccessToken, error)
	}
	createToken struct{}
)

func NewCreateToken() CreateToken {
	return createToken{}
}

func (u createToken) Do() (d.AccessToken, error) {
	result, err := u.create()
	if err != nil {
		return d.AccessToken{}, err
	}

	return result, nil
}

func (u createToken) create() (d.AccessToken, error) {
	claims := myauth.AuthClaims{StandardClaims: &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(config.Cfg.Auth.TokenExpiration)).Unix(),
		Issuer:    "zo.auth.service",
	},
		TokenType: "accessToken",
	}

	// https://pkg.go.dev/github.com/golang-jwt/jwt#NewWithClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	j, err := token.SignedString([]byte(config.Cfg.Auth.SignKey))
	log.Printf("signed: %v", j)

	if err != nil {
		return d.AccessToken{}, err
	}

	return d.AccessToken{Jwt: j, ExpiresAt: claims.ExpiresAt}, nil
}
