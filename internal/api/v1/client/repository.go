package client

import (
	"context"
	"database/sql"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
)

type clientRepository struct {
}

func (r *clientRepository) Find(ctx context.Context, id int, secret string) (*Client, error) {
	m := Client{}
	err := mydb.Db.QueryRowContext(ctx, "SELECT * FROM Clients WHERE id = ? AND secret = ?",
		id, secret).
		Scan(
			&m.Id,
			&m.Secret,
			&m.CreatedAt,
			&m.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &m, nil
}
