package zo

import (
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform"
)

type GetAllResponse struct {
	platform.ResponseBase
	Zos []Zo `json:"zos"`
}

type GetResponse struct {
	platform.ResponseBase
	Zo *Zo `json:"zo"`
}

type PostResponse struct {
	platform.ResponseBase
	Zo *Zo `json:"zo"`
}

type PutResponse struct {
	platform.ResponseBase
	Zo *Zo `json:"zo"`
}

type DeleteResponse struct {
	platform.ResponseBase
}
