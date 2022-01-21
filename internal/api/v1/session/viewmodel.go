package session

import (
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
)

type PostResponse struct {
	*myhttp.ResponseBase
	*SessionData
}

func NewPostResponse(res *myhttp.ResponseBase, sessionData *SessionData) *PostResponse {
	return &PostResponse{ResponseBase: res, SessionData: sessionData}
}
