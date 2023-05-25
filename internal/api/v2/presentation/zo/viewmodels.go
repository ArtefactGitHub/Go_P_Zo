package zo

import (
	"database/sql"
	"time"

	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
)

type (
	PostRequest struct {
		AchievementDate time.Time `json:"achievementdate"`
		Exp             int       `json:"exp"`
		CategoryId      int       `json:"category_id"`
		Message         string    `json:"message"`
	}
	PostResponse struct {
		*myhttp.ResponseBase
		Zo Zo `json:"zo"`
	}
	SimpleResponse struct {
		*myhttp.ResponseBase
	}

	Zo struct {
		Id              int          `json:"id"`
		AchievementDate time.Time    `json:"achievementdate"`
		Exp             int          `json:"exp"`
		CategoryId      int          `json:"category_id"`
		Message         string       `json:"message"`
		CreatedAt       time.Time    `json:"createdat"`
		UpdatedAt       sql.NullTime `json:"updatedat"`
		UserId          int          `json:"user_id"`
	}
)

func NewSimpleResponse(res *myhttp.ResponseBase) *SimpleResponse {
	return &SimpleResponse{ResponseBase: res}
}

func NewZoResponse(res *myhttp.ResponseBase, zo d.Zo) *PostResponse {
	z := ToResponse(zo)
	return &PostResponse{ResponseBase: res, Zo: z}
}

func ToResponse(zo d.Zo) Zo {
	return Zo{
		Id:              zo.Id,
		AchievementDate: zo.AchievementDate,
		Exp:             zo.Exp,
		CategoryId:      zo.CategoryId,
		Message:         zo.Message,
		CreatedAt:       zo.CreatedAt,
		UpdatedAt:       zo.UpdatedAt,
		UserId:          zo.UserId,
	}
}
