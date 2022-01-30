package user

import (
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
)

type GetAllResponse struct {
	myhttp.ResponseBase
	Users []User `json:"zos"`
}

type GetResponse struct {
	myhttp.ResponseBase
	User *User `json:"user"`
}

type PostResponse struct {
	myhttp.ResponseBase
	User *User `json:"user"`
}

type PutResponse struct {
	myhttp.ResponseBase
	User *User `json:"user"`
}

type DeleteResponse struct {
	myhttp.ResponseBase
}

type UserTokenRequest struct {
	Identifier string `json:"identifier"`
	Secret     string `json:"secret"`
}

func NewUserTokenRequest(
	identifier string,
	secret string,
) UserTokenRequest {
	return UserTokenRequest{
		Identifier: identifier,
		Secret:     secret}
}

type PostUserTokenResponse struct {
	*myhttp.ResponseBase
	*UserToken
}

func NewPostUserTokenResponse(res *myhttp.ResponseBase, usertoken *UserToken) *PostUserTokenResponse {
	return &PostUserTokenResponse{ResponseBase: res, UserToken: usertoken}
}
