package myauth

import (
	"github.com/golang-jwt/jwt"
)

type AuthClaims struct {
	*jwt.StandardClaims
	TokenType string
}
