package session

import (
	"context"
)

type Repository interface {
	Find(ctx context.Context, identifier string, password string) (SessionData, error)
}
