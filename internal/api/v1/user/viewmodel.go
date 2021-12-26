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
