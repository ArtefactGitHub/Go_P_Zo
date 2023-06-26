package zo

import (
	"context"

	domain "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
)

type (
	Create interface {
		Do(context.Context, domain.Zo) (domain.Zo, error)
	}
	create struct {
		r domain.Repository
	}
)

func NewCreate(r domain.Repository) Create {
	return create{r: r}
}

func (u create) Do(ctx context.Context, zo domain.Zo) (domain.Zo, error) {
	_, err := u.r.Create(ctx, zo)
	if err != nil {
		return nil, err
	}

	return zo, nil
}
