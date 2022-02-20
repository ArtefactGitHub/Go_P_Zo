package user

import (
	"fmt"

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

func NewResponseUsers(models []User) []responseUser {
	result := []responseUser{}
	for _, m := range models {
		result = append(result, responseUser{Id: m.Id, GivenName: m.GivenName, FamilyName: m.FamilyName, Email: m.Email})
	}
	return result
}

func (u *responseUser) FullName() string {
	return fmt.Sprintf("%s %s", u.FamilyName, u.GivenName)
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
	User      *responseUser `json:"user"`
	UserToken *UserToken    `json:"usertoken"`
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

// UserCategory
type responseUserCategory struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	ColorId    int    `json:"color_id"`
	CreateType int    `json:"create_type"`
	UserId     int    `json:"user_id"`
}

func NewResponseUserCategory(userId int, givenName, familyName, email string) *responseUser {
	return &responseUser{Id: userId, GivenName: givenName, FamilyName: familyName, Email: email}
}

func NewResponseUserCategories(models []Category) []responseUserCategory {
	result := []responseUserCategory{}
	for _, m := range models {
		result = append(result, responseUserCategory{
			Id: m.Id, Name: m.Name, ColorId: m.ColorId, CreateType: m.CreateType, UserId: m.UserId})
	}
	return result
}

type GetAllUserCategoryResponse struct {
	*myhttp.ResponseBase
	Categories []responseUserCategory `json:"categories"`
}

func NewGetAllUserCategoryResponse(res *myhttp.ResponseBase, categories []responseUserCategory) *GetAllUserCategoryResponse {
	return &GetAllUserCategoryResponse{ResponseBase: res, Categories: categories}
}
