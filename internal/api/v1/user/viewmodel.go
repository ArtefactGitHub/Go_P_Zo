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

func NewResponseUser(userId int, givenName, familyName, email string) *responseUser {
	return &responseUser{Id: userId, GivenName: givenName, FamilyName: familyName, Email: email}
}

func NewResponseUsers(users []User) []responseUser {
	result := []responseUser{}
	for _, u := range users {
		result = append(result, responseUser{Id: u.Id, GivenName: u.GivenName, FamilyName: u.FamilyName, Email: u.Email})
	}
	return result
}

type GetAllResponse struct {
	myhttp.ResponseBase
	Users []responseUser `json:"users"`
}

type GetResponse struct {
	myhttp.ResponseBase
	User *responseUser `json:"user"`
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
