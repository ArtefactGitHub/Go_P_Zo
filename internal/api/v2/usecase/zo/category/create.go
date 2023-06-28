package category

import (
	"context"

	domain "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
)

type (
	Create interface {
		Do(context.Context, domain.Category) (domain.Category, error)
	}
	create struct {
		r domain.CategoryRepository
	}
)

func NewCreate(r domain.CategoryRepository) Create {
	return create{r: r}
}

func (u create) Do(ctx context.Context, v domain.Category) (domain.Category, error) {
	_, err := u.r.Create(ctx, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
