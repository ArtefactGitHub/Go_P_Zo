package session

import (
	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/session"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
)

type (
	PostRequest struct {
		Identifier string `json:"identifier"`
		Secret     string `json:"secret"`
	}
	PostResponse struct {
		*myhttp.ResponseBase
		SessionData *SessionData
	}
	SessionData struct {
		GivenName  string `json:"given_name"`
		FamilyName string `json:"family_name"`
		Email      string `json:"email"`
	}
)

func NewPostResponse(res *myhttp.ResponseBase, data d.SessionData) *PostResponse {
	return &PostResponse{ResponseBase: res, SessionData: &SessionData{
		GivenName:  data.GivenName,
		FamilyName: data.FamilyName,
		Email:      data.Email,
	}}
}
