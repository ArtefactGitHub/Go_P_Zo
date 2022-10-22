package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mycontext"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
)

const (
	AuthTokenSplit = "Bearer "
)

type JwtMiddleware struct {
	next    IMiddleware
	authKey string
}

func NewJwtMiddleware(config *config.Config) (IMiddleware, error) {
	switch {
	case config == nil:
		return nil, errors.New("config not found")
	case config.Auth.SignKey == "":
		return nil, errors.New("config.SignKey not found")
	default:
		return &JwtMiddleware{authKey: config.Auth.SignKey}, nil
	}
}

func (m *JwtMiddleware) SetNext(next IMiddleware) {
	m.next = next
}

func (m *JwtMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	tokenHeader := req.Header.Get(myhttp.AuthTokenHeaderName)

	if tokenHeader != "" {
		log.Printf("tokenHeader: %s", tokenHeader)

		tokenString, err := getExtractedJwtToken(tokenHeader)
		log.Printf("tokenString: %s", tokenHeader)
		if err != nil {
			myhttp.WriteError(w, err, http.StatusUnauthorized, "")
			return
		}

		token, err := parseJwtToken(tokenString, m.authKey)
		log.Printf("token: %v", token)
		if err != nil {
			myhttp.WriteError(w, err, http.StatusUnauthorized, "parse token failure")
			return
		}

		if _, err := m.verifyToken(token); err != nil {
			myhttp.WriteError(w, err, http.StatusUnauthorized, "verify token failure")
			return
		}

		// set value to context
		ctx = mycontext.NewContext(ctx, mycontext.AuthorizedKey, true)
	}

	if m.next != nil {
		m.next.ServeHTTP(w, req.WithContext(ctx))
	}
}

func (m *JwtMiddleware) verifyToken(token *jwt.Token) (*myauth.AuthClaims, error) {
	if claims, ok := token.Claims.(*myauth.AuthClaims); ok {
		log.Printf("claims: %v", claims)
		if claims.Issuer != myauth.Issuer {
			return nil, errors.New("invalid issuer")
		}
		now := time.Now().Unix()
		if claims.ExpiresAt < now {
			return nil, errors.New("token expired")
		}

		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

func getExtractedJwtToken(tokenHeader string) (string, error) {
	extractedToken := strings.Split(tokenHeader, AuthTokenSplit)
	if len(extractedToken) == 2 {
		return strings.TrimSpace(extractedToken[1]), nil
	}

	return "", errors.New("invalid token")
}

func parseJwtToken(tokenString string, authKey string) (*jwt.Token, error) {
	// parse the jwt.Token
	// https://pkg.go.dev/github.com/golang-jwt/jwt/v4#Parse
	token, err := jwt.ParseWithClaims(tokenString, &myauth.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(authKey), nil
	})

	if token.Valid {
		return token, nil
	}
	if ve, ok := err.(*jwt.ValidationError); ok {
		switch {
		case ve.Errors&jwt.ValidationErrorMalformed != 0:
			return nil, fmt.Errorf("token is malformed. %v", token)
		case ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0:
			return nil, fmt.Errorf("token is either expired or not active yet. %v", token)
		default:
			return nil, fmt.Errorf("couldn't handle this token. %s", err)
		}
	} else {
		return nil, fmt.Errorf("couldn't handle this token. %s", err)
	}
}
