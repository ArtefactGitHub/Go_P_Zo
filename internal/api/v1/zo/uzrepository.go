package zo

import (
	"context"
	"database/sql"
	"time"
)

// TODO：仮実装

type UserZos struct {
	UserId    int          `json:"user_id"`
	ZoId      int          `json:"zo_id"`
	CreatedAt time.Time    `json:"createdat"`
	UpdatedAt sql.NullTime `json:"updatedat"`
}
type UserZosRepository struct {
}

func (r *UserZosRepository) Findall(ctx context.Context) ([]UserZos, error) {
	return nil, nil
}

func (r *UserZosRepository) Find(ctx context.Context, id int) (*UserZos, error) {
	return nil, nil
}

func (r *UserZosRepository) Create(ctx context.Context, uz *UserZos) (int, error) {
	return -1, nil
}

func (r *UserZosRepository) CreateTx(ctx context.Context, tx *sql.Tx, uz *UserZos) (int, error) {
	return -1, nil
}

func (r *UserZosRepository) Update(ctx context.Context, uz *UserZos) error {
	return nil
}

func (r *UserZosRepository) Delete(ctx context.Context, id int) error {
	return nil
}
