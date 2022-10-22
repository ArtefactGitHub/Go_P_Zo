package client

import (
	"context"
)

type Repository interface {
	Find(ctx context.Context, id int, secret string) (Client, error)
}
