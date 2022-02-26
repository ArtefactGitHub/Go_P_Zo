package myauth

import (
	"errors"
	"log"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/golang-jwt/jwt"
)

type AuthClaims struct {
	*jwt.StandardClaims
	TokenType string
}

type UserTokenClaims struct {
	*jwt.StandardClaims
	TokenType string
	UserId    int
}

func CreateUserTokenJwt(userId int, expiredAt time.Time) (string, error) {
	claims := UserTokenClaims{StandardClaims: &jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Issuer:    "zo.auth.service",
	},
		TokenType: "userToken",
		UserId:    userId,
	}

	// https://pkg.go.dev/github.com/golang-jwt/jwt#NewWithClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := token.SignedString([]byte(config.Cfg.Auth.SignKey))
	log.Printf("signed usertoken: %v", jwt)

	if err != nil {
		return "", err
	}

	return jwt, nil
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
