package zo

import (
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
)

type GetAllResponse struct {
	myhttp.ResponseBase
	Zos []Zo `json:"zos"`
}

type GetResponse struct {
	myhttp.ResponseBase
	Zo *Zo `json:"zo"`
}

type PostResponse struct {
	myhttp.ResponseBase
	Zo *Zo `json:"zo"`
}

type PutResponse struct {
	myhttp.ResponseBase
	Zo *Zo `json:"zo"`
}

type DeleteResponse struct {
	myhttp.ResponseBase
}
