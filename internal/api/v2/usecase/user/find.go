package user

import (
	"context"

	domain "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/user"
)

type (
	Find interface {
		Do(context.Context, int) (domain.User, error)
	}
	find struct {
		r domain.Repository
	}
)

func NewFind(r domain.Repository) Find {
	return find{r: r}
}

func (u find) Do(ctx context.Context, id int) (domain.User, error) {
	v, err := u.r.Find(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return v, nil
}
