package zo

import (
	"context"
	"database/sql"
)

// TODO：仮実装

type UserZosRepository struct {
}

func (r *UserZosRepository) Findall(ctx context.Context) ([]UserZo, error) {
	return nil, nil
}

func (r *UserZosRepository) Find(ctx context.Context, id int) (*UserZo, error) {
	return nil, nil
}

func (r *UserZosRepository) Create(ctx context.Context, uz *UserZo) (int, error) {
	return -1, nil
}

func (r *UserZosRepository) CreateTx(ctx context.Context, tx *sql.Tx, uz *UserZo) (int, error) {
	return -1, nil
}

func (r *UserZosRepository) Update(ctx context.Context, uz *UserZo) error {
	return nil
}

func (r *UserZosRepository) Delete(ctx context.Context, id int) error {
	return nil
}
