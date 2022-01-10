package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mycontext"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
)

const (
	AuthTokenHeaderName = "Authorization"
	AuthTokenSplit      = "Bearer "
)

type JwtMiddleware struct {
	next    IMiddleware
	authKey string
}

func NewJwtMiddleware(config *config.Config) (IMiddleware, error) {
	if config.Auth.SignKey == "" {
		return nil, errors.New("config.SignKey not found")
	} else {
		// authtest(config)
		return &JwtMiddleware{authKey: config.Auth.SignKey}, nil
	}
}

func (m *JwtMiddleware) SetNext(next IMiddleware) {
	m.next = next
}

func (m *JwtMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	tokenHeader := req.Header.Get(AuthTokenHeaderName)
	if tokenHeader != "" {
		tokenString, err := getExtractedToken(tokenHeader)
		if err != nil {
			myhttp.WriteError(w, err, http.StatusBadRequest, "")
			return
		}

		token, err := validateToken(tokenString, m.authKey)
		if err != nil {
			myhttp.WriteError(w, err, http.StatusBadRequest, "parse token failure")
			return
		}

		if _, ok := token.Claims.(*myauth.AuthClaims); !ok {
			// set value to context
			ctx = mycontext.NewContext(ctx, mycontext.AuthorizedKey, true)
		} else {
			myhttp.WriteError(w, errors.New("token.Claims is invalid"), http.StatusBadRequest, "")
			return
		}
	}

	if m.next != nil {
		m.next.ServeHTTP(w, req.WithContext(ctx))
	}
}

func getExtractedToken(tokenHeader string) (string, error) {
	extractedToken := strings.Split(tokenHeader, AuthTokenSplit)
	if len(extractedToken) == 2 {
		return strings.TrimSpace(extractedToken[1]), nil
	}

	return "", errors.New("invalid token")
}

func validateToken(tokenString string, authKey string) (*jwt.Token, error) {
	// Parse the token
	// https://pkg.go.dev/github.com/golang-jwt/jwt/v4#Parse
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(authKey), nil
	})

	if token.Valid {
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, fmt.Errorf("token is malformed. %v", token)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, fmt.Errorf("token is either expired or not active yet. %v", token)
		} else {
			return nil, fmt.Errorf("couldn't handle this token. %s", err)
		}
	} else {
		return nil, fmt.Errorf("couldn't handle this token. %s", err)
	}
	return token, nil
}
