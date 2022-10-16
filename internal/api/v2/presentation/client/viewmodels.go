package client

import (
	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/client"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
)

type (
	PostRequest struct {
		Id     int    `json:"id"`
		Secret string `json:"secret"`
	}
	PostResponse struct {
		*myhttp.ResponseBase
		AccessToken AccessToken `json:"access_token"`
	}

	AccessToken struct {
		Jwt       string `json:"jwt"`
		ExpiresAt int64  `json:"expires_at"`
	}
)

func NewPostResponse(res *myhttp.ResponseBase, token d.AccessToken) *PostResponse {
	t := AccessToken{
		Jwt:       token.Jwt(),
		ExpiresAt: token.ExpiresAt(),
	}
	return &PostResponse{ResponseBase: res, AccessToken: t}
}
