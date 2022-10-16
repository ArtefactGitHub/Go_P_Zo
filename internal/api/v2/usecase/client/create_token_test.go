package client_test

import (
	. "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/client"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
	"github.com/golang-jwt/jwt"
	"testing"
	"time"
)

func Test_createToken_Do(t *testing.T) {
	tests := []struct {
		name          string
		wantTokenType string
		wantIssuer    string
		wantErr       bool
		wantExpiresAt time.Time
	}{
		{
			name:          "success",
			wantTokenType: "accessToken",
			wantIssuer:    "zo.auth.service",
			wantExpiresAt: time.Now().Add(time.Minute * time.Duration(config.Cfg.Auth.TokenExpiration)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewCreateToken()
			got, err := u.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tokenString := got.Jwt()
			claims := myauth.AuthClaims{}
			_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(config.Cfg.Auth.SignKey), nil
			})
			if err != nil {
				t.Errorf(err.Error())
			}

			if claims.TokenType != "accessToken" {
				t.Errorf("invalid TokenType got = %v, want %v", claims.TokenType, tt.wantTokenType)
			}
			if claims.Issuer != "zo.auth.service" {
				t.Errorf("invalid Issuer got = %v, want %v", claims.TokenType, tt.wantIssuer)
			}
			expiresAtSub := int64(tt.wantExpiresAt.Sub(time.Unix(claims.ExpiresAt, 0)) / time.Second)
			if expiresAtSub != 0 {
				t.Errorf("invalid expiresAt got = %v, want %v", claims.ExpiresAt, tt.wantExpiresAt)
			}
		})
	}
}
