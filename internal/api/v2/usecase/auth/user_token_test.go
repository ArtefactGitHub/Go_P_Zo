package auth_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/auth"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/user"
	i "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure"
	ia "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/auth"
	iu "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure/user"
	. "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/usecase/auth"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_create_Do(t *testing.T) {
	type fields struct {
		r  auth.Repository
		ur user.Repository
	}
	type args struct {
		ctx  context.Context
		data CreateTokenData
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          auth.UserToken
		wantTokenType string
		wantIssuer    string
		wantUserId    int
		wantExpiresAt time.Time
		wantErr       bool
	}{
		{
			name: "正常系：認証が成功し、想定通りのトークン情報が成功している",
			fields: fields{
				r:  ia.NewRepository(),
				ur: iu.NewRepository(),
			},
			args: args{
				data: CreateTokenData{
					Identifier: "test@com",
					Secret:     "password",
				},
			},
			want: auth.NewUserToken(
				10003,
				1,
				"token",
				time.Now(),
				time.Now(),
				sql.NullTime{}),
			wantTokenType: myauth.UserTokenType,
			wantIssuer:    myauth.Issuer,
			wantExpiresAt: time.Now().Add(time.Minute * 1200),
			wantUserId:    1,
			wantErr:       false,
		},
		{
			name: "異常系：認証が失敗(identifier が不正)",
			fields: fields{
				r:  ia.NewRepository(),
				ur: iu.NewRepository(),
			},
			args: args{
				data: CreateTokenData{
					Identifier: "dummy@com",
					Secret:     "password",
				},
			},
			wantErr: true,
		},
		{
			name: "異常系：認証が失敗(Secret が不正)",
			fields: fields{
				r:  ia.NewRepository(),
				ur: iu.NewRepository(),
			},
			args: args{
				data: CreateTokenData{
					Identifier: "test@com",
					Secret:     "dummy",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewCreate(tt.fields.r, tt.fields.ur)
			c := context.WithValue(context.Background(), i.KeyDB, DB)
			c = context.WithValue(c, i.KeyTX, TX)

			got, err := u.Do(c, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(auth.UserToken{}, "Token", "ExpiredAt", "CreatedAt", "UpdatedAt")); diff != "" {
				t.Errorf("u.Do() value is mismatch: %s", diff)
			}

			// claimsの検証
			claims, err := myauth.CreateUserTokenClaims(got.Token)
			if err != nil {
				t.Errorf("CreateUserTokenClaims() error = %v", err)
				return
			}
			if claims.TokenType != tt.wantTokenType {
				t.Errorf("invalid TokenType got = %v, want %v", claims.TokenType, tt.wantTokenType)
			}
			if claims.Issuer != tt.wantIssuer {
				t.Errorf("invalid Issuer got = %v, want %v", claims.TokenType, tt.wantIssuer)
			}
			expiresAtSub := int64(tt.wantExpiresAt.Sub(time.Unix(claims.ExpiresAt, 0)) / time.Second)
			if expiresAtSub != 0 {
				t.Errorf("invalid expiresAt got = %v, want %v", claims.ExpiresAt, tt.wantExpiresAt.Unix())
			}
			if claims.UserId != tt.wantUserId {
				t.Errorf("invalid UserId got = %v, want %v", claims.UserId, tt.wantUserId)
			}
		})
	}
}
