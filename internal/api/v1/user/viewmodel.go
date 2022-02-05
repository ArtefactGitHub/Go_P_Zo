package user

import (
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
)

type responseUser struct {
	Id         int    `json:"id"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Email      string `json:"email"`
}

type GetAllResponse struct {
	myhttp.ResponseBase
	Users []User `json:"zos"`
}

type GetResponse struct {
	myhttp.ResponseBase
	User *responseUser `json:"user"`
}

func NewResponseUser(userId int, givenName, familyName, email string) *responseUser {
	return &responseUser{Id: userId, GivenName: givenName, Email: email}
}

type PostResponse struct {
	myhttp.ResponseBase
	User *responseUser `json:"user"`
}

type PutResponse struct {
	myhttp.ResponseBase
	User *responseUser `json:"user"`
}

type DeleteResponse struct {
	myhttp.ResponseBase
}

// UserToken
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
