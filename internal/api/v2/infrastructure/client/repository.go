package client

import (
	"context"
	"database/sql"
	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/client"
	e "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	"time"
)

type repository struct {
}

func NewRepository() d.Repository {
	return repository{}
}

func (r repository) Find(ctx context.Context, id int, secret string) (d.Client, error) {
	var createdAt time.Time
	var updatedAt sql.NullTime
	err := mydb.Db.QueryRowContext(ctx, "SELECT * FROM Clients WHERE id = ? AND secret = ?",
		id, secret).
		Scan(
			&id,
			&secret,
			&createdAt,
			&updatedAt)
	if err == sql.ErrNoRows {
		return nil, e.NotFound
	} else if err != nil {
		return nil, err
	}

	return d.NewClient(id, secret, createdAt, updatedAt), nil
}
