package user

import (
	"context"
)

type (
	Repository interface {
		FindByIdentifier(ctx context.Context, identifier string, password string) (User, error)
		Find(ctx context.Context, id int) (User, error)
		Create(ctx context.Context, u User) (User, error)
	}
)
