package client

import (
	domain "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/client"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
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
	expiredAt := time.Now().Add(time.Minute * time.Duration(config.Cfg.Auth.TokenExpiration))
	jwt, err := myauth.CreateAccessTokenJwt(expiredAt)
	if err != nil {
		return nil, err
	}

	return domain.NewAccessToken(jwt, expiredAt.Unix()), nil
}
