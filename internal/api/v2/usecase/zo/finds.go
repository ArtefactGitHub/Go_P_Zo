package zo

import (
	"context"

	derr "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
	domain "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
)

type (
	Finds interface {
		Do(context.Context, int) ([]domain.Zo, error)
	}
	finds struct {
		r domain.Repository
	}
)

func NewFinds(r domain.Repository) Finds {
	return finds{r: r}
}

func (u finds) Do(ctx context.Context, id int) ([]domain.Zo, error) {
	v, err := u.r.Finds(ctx, id)
	if err != nil {
		return nil, err
	} else if len(v) == 0 {
		return nil, derr.NotFound
	}

	return v, nil
}
