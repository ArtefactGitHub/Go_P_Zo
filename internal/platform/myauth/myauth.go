package myauth

import (
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
