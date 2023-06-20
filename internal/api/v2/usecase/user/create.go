package user

import (
	"context"
	"errors"
	"fmt"

	derr "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
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
	_, err := u.r.FindByIdentifier(ctx, user.Email, user.Password)
	switch {
	case err == nil:
		return domain.User{}, fmt.Errorf("%s: %w", "既に存在するメールアドレスです", derr.Conflict)
	case !errors.Is(derr.NotFound, err):
		return domain.User{}, fmt.Errorf("%s: %w", err.Error(), derr.BadRequest)
	}

	result, err := u.r.Create(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return result, nil
}
