package auth

import (
	"context"
)

type Repository interface {
	Find(ctx context.Context, id int) (UserToken, error)
	Create(ctx context.Context, m UserToken) (UserToken, error)
}
