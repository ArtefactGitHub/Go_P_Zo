package zo

import (
	"github.com/ArtefactGitHub/Go_P_Zo/internal/models"
)

type GetAllResponse struct {
	models.ResponseBase
	Zos []Zo `json:"zos"`
}

type GetResponse struct {
	models.ResponseBase
	Zo *Zo `json:"zo"`
}

type PostResponse struct {
	models.ResponseBase
	Zo *Zo `json:"zo"`
}

type PutResponse struct {
	models.ResponseBase
	Zo *Zo `json:"zo"`
}

type DeleteResponse struct {
	models.ResponseBase
}
