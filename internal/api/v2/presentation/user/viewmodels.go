package user

import (
	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/user"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
)

type (
	PostResponse struct {
		*myhttp.ResponseBase
		User User `json:"user"`
	}
	SimpleResponse struct {
		*myhttp.ResponseBase
	}

	User struct {
		Id         int    `json:"id"`
		GivenName  string `json:"given_name"`
		FamilyName string `json:"family_name"`
		Email      string `json:"email"`
	}
)

func NewSimpleResponse(res *myhttp.ResponseBase) *SimpleResponse {
	return &SimpleResponse{ResponseBase: res}
}

func NewResponse(res *myhttp.ResponseBase, user d.User) *PostResponse {
	return &PostResponse{ResponseBase: res, User: ToResponse(user)}
}

func ToResponse(v d.User) User {
	return User{
		v.Id,
		v.GivenName,
		v.FamilyName,
		v.Email,
	}
}
