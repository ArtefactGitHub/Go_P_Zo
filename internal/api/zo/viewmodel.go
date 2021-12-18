package zo

import (
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
)

type getAllResponse struct {
	myhttp.ResponseBase
	Zos []zo `json:"zos"`
}

type getResponse struct {
	myhttp.ResponseBase
	Zo *zo `json:"zo"`
}

type postResponse struct {
	myhttp.ResponseBase
	Zo *zo `json:"zo"`
}

type putResponse struct {
	myhttp.ResponseBase
	Zo *zo `json:"zo"`
}

type deleteResponse struct {
	myhttp.ResponseBase
}
