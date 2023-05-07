package myauth_test

import (
	"errors"
	"testing"
	"time"

	. "github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"

	"github.com/golang-jwt/jwt"
	"github.com/google/go-cmp/cmp"
)

func TestCreateUserTokenClaims(t *testing.T) {
	userID := 1
	expiration := time.Now().Add(time.Hour * 24)
	tokenString, _ := CreateUserTokenJwt(userID, expiration)

	tests := []struct {
		name        string
		userToken   string
		expected    *UserTokenClaims
		expectedErr error
	}{
		{
			name:      "【正常系】指定のUserIDや有効期限に加え、IssuerやTokenTypeが正しくセットされる",
			userToken: tokenString,
			expected: &UserTokenClaims{
				UserId: userID,
				StandardClaims: &jwt.StandardClaims{
					ExpiresAt: expiration.Unix(),
					Issuer:    Issuer,
				},
				TokenType: UserTokenType,
			},
			expectedErr: nil,
		},
		{
			name:        "【異常系】空のトークンを渡した場合",
			userToken:   "",
			expected:    nil,
			expectedErr: errors.New("invalid userToken"),
		},
		{
			name:        "【異常系】無効なトークンを渡した場合",
			userToken:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2MTg3MTQxOTF9.WIw7zP4l4pJ7D1A39o0XzJt1bGt_mKe8ZTxhYw1JYbk",
			expected:    nil,
			expectedErr: errors.New("can not parse userToken"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := CreateUserTokenClaims(test.userToken)

			if diff := cmp.Diff(test.expectedErr, err, cmp.Comparer(func(x, y error) bool { return x == nil && y == nil || x.Error() == y.Error() })); diff != "" {
				t.Errorf("CreateUserTokenClaims error mismatch (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(test.expected, result, cmp.AllowUnexported(UserTokenClaims{}, jwt.StandardClaims{})); diff != "" {
				t.Errorf("CreateUserTokenClaims result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
