package zo

import (
	"context"

	domain "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
)

type (
	Find interface {
		Do(context.Context, int) (domain.Zo, error)
	}
	find struct {
		r domain.Repository
	}
)

func NewFind(r domain.Repository) Find {
	return find{r: r}
}

func (u find) Do(ctx context.Context, id int) (domain.Zo, error) {
	v, err := u.r.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return v, nil
}
