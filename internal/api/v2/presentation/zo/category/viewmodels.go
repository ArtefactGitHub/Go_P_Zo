package category

import (
	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
)

type (
	PostRequest struct {
		Name string `json:"name"`
	}
	PostResponse struct {
		*myhttp.ResponseBase
		Category Category `json:"category"`
	}
	CategoriesResponse struct {
		*myhttp.ResponseBase
		Categories []Category `json:"categories"`
	}
	SimpleResponse struct {
		*myhttp.ResponseBase
	}

	Category struct {
		ID     int    `json:"id"`
		UserID int    `json:"user_id"`
		Name   string `json:"name"`
	}
)

func NewResponse(res *myhttp.ResponseBase, v d.Category) *PostResponse {
	c := ToResponse(v)
	return &PostResponse{ResponseBase: res, Category: c}
}

func NewCategoriesResponse(res *myhttp.ResponseBase, categories []d.Category) *CategoriesResponse {
	cs := []Category{}
	for _, v := range categories {
		cs = append(cs, ToResponse(v))

	}
	return &CategoriesResponse{ResponseBase: res, Categories: cs}
}

func ToResponse(v d.Category) Category {
	return Category{
		ID:     v.ID(),
		Name:   v.Name(),
		UserID: v.UserID(),
	}
}
