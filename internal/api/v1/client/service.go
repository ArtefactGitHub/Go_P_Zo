package client

import (
	"context"
	"log"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
	"github.com/golang-jwt/jwt"
)

type clientService struct {
	r clientRepository
}

func (s *clientService) Exist(ctx context.Context, id int, secret string) (bool, error) {
	result, err := s.r.Find(ctx, id, secret)
	if err != nil {
		return false, err
	}

	return result != nil, nil
}

func (s *clientService) find(ctx context.Context, id int, secret string) (*Client, error) {
	result, err := s.r.Find(ctx, id, secret)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *clientService) createAccessToken() (*AccessToken, error) {
	claims := myauth.AuthClaims{StandardClaims: &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(config.Cfg.Auth.TokenExpiration)).Unix(),
		Issuer:    "zo.auth.service",
	},
		TokenType: "accessToken",
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
