package auth

import (
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/auth"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
)

type (
	PostResponse struct {
		*myhttp.ResponseBase
		UserToken *UserToken `json:"user_token,omitempty"`
	}

	UserToken struct {
		UserId    int       `json:"user_id"`
		Token     string    `json:"token"`
		ExpiredAt time.Time `json:"expiredat"`
	}
)

func NewPostResponse(res *myhttp.ResponseBase, token *auth.UserToken) *PostResponse {
	if token == nil {
		return &PostResponse{ResponseBase: res, UserToken: nil}
	}
	t := &UserToken{
		UserId:    token.UserId,
		Token:     token.Token,
		ExpiredAt: token.ExpiredAt,
	}
	return &PostResponse{ResponseBase: res, UserToken: t}
}
