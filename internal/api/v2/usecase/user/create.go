package user

import (
	"context"

	domain "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/user"
)

type (
	Create interface {
		Do(context.Context, domain.User) (domain.User, error)
	}
	create struct {
		r domain.Repository
	}
)

func NewCreate(r domain.Repository) Create {
	return create{r: r}
}

func (u create) Do(ctx context.Context, user domain.User) (domain.User, error) {
	_, err := u.r.Create(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
