package zo

import (
	"context"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Zo, error)
	FindAllByUserId(ctx context.Context, userId int) ([]Zo, error)
	Find(ctx context.Context, id int) (Zo, error)
	Create(ctx context.Context, z Zo) (int, error)
	Update(ctx context.Context, z Zo) error
	Delete(ctx context.Context, id int) error
}
