package category

import (
	"context"
	"fmt"

	domain "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
)

type (
	Finds interface {
		Do(context.Context, int) ([]domain.Category, error)
	}
	finds struct {
		r domain.CategoryRepository
	}
)

func NewFinds(r domain.CategoryRepository) Finds {
	return finds{r: r}
}

func (u finds) Do(ctx context.Context, id int) ([]domain.Category, error) {
	fmt.Printf("id: %d \n", id)
	v, err := u.r.Finds(ctx, id)
	if err != nil {
		return nil, err
	}

	return v, nil
}
