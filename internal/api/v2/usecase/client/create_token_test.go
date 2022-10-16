package client_test

import (
	. "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/client"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
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
			wantTokenType: myauth.TokenType,
			wantIssuer:    myauth.Issuer,
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

			claims, err := myauth.CreateUserTokenClaims(got.Jwt())
			if err != nil {
				t.Fatalf(err.Error())
			}

			if claims.TokenType != myauth.TokenType {
				t.Errorf("invalid TokenType got = %v, want %v", claims.TokenType, tt.wantTokenType)
			}
			if claims.Issuer != myauth.Issuer {
				t.Errorf("invalid Issuer got = %v, want %v", claims.TokenType, tt.wantIssuer)
			}
			expiresAtSub := int64(tt.wantExpiresAt.Sub(time.Unix(claims.ExpiresAt, 0)) / time.Second)
			if expiresAtSub != 0 {
				t.Errorf("invalid expiresAt got = %v, want %v", claims.ExpiresAt, tt.wantExpiresAt)
			}
		})
	}
}
