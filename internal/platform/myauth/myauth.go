package myauth

import (
	"errors"
	"log"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/golang-jwt/jwt"
)

const (
	Issuer        = "zo.auth.service"
	TokenType     = "accessToken"
	UserTokenType = "userToken"
)

type (
	AuthClaims struct {
		*jwt.StandardClaims
		TokenType string
	}
	UserTokenClaims struct {
		*jwt.StandardClaims
		TokenType string
		UserId    int
	}
)

func CreateAccessTokenJwt(expiredAt time.Time) (string, error) {
	claims := AuthClaims{StandardClaims: &jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Issuer:    Issuer,
	},
		TokenType: TokenType,
	}

	// https://pkg.go.dev/github.com/golang-jwt/jwt#NewWithClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	j, err := token.SignedString([]byte(config.Cfg.Auth.SignKey))
	if err != nil {
		return "", err
	}

	return j, nil
}

func CreateUserTokenJwt(userId int, expiredAt time.Time) (string, error) {
	claims := UserTokenClaims{StandardClaims: &jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Issuer:    Issuer,
	},
		TokenType: UserTokenType,
		UserId:    userId,
	}

	// https://pkg.go.dev/github.com/golang-jwt/jwt#NewWithClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	j, err := token.SignedString([]byte(config.Cfg.Auth.SignKey))

	if err != nil {
		return "", err
	}

	return j, nil
}

func CreateUserTokenClaims(userToken string) (*UserTokenClaims, error) {
	if userToken == "" {
		return nil, errors.New("invalid userToken")
	}

	token, err := jwt.ParseWithClaims(userToken, &UserTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.Auth.SignKey), nil
	})
	if err != nil {
		return nil, errors.New("can not parse userToken")
	}

	if claims, ok := token.Claims.(*UserTokenClaims); ok && token.Valid {
		log.Printf("token: %v", claims)
		return claims, nil
	}
	return nil, nil
}
