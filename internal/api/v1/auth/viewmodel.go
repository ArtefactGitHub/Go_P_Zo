package auth

import (
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
)

type PostResponse struct {
	myhttp.ResponseBase
	AccessToken AccessToken `json:"access_token"`
}

type AccessToken struct {
	Jwt       string `json:"jwt"`
	ExpiresAt int64  `json:"expires_at"`
}
