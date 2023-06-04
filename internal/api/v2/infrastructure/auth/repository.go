package auth

import (
	"context"
	"database/sql"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/auth"
	e "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
)

type repository struct {
}

func NewRepository() auth.Repository {
	return &repository{}
}

func (r *repository) Create(ctx context.Context, m auth.UserToken) (int, error) {
	result, err := mydb.Db.ExecContext(ctx, `
		INSERT INTO UserTokens(id, user_id, token, expiredAt, createdAt, updatedAt)
			VALUES(?,?,?,?,?,?)`,
		nil,
		&m.UserId,
		&m.Token,
		&m.ExpiredAt,
		&m.CreatedAt,
		&m.UpdatedAt)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	m.Id = int(id)
	return m.Id, nil
}

func (r repository) Find(ctx context.Context, id int) (auth.UserToken, error) {
	result := auth.UserToken{}
	err := mydb.Db.QueryRowContext(ctx, "SELECT * FROM UserTokens WHERE id = ?",
		id).
		Scan(
			&result.Id,
			&result.UserId,
			&result.Token,
			&result.ExpiredAt,
			&result.CreatedAt,
			&result.UpdatedAt)
	if err == sql.ErrNoRows {
		return auth.UserToken{}, e.NotFound
	} else if err != nil {
		return auth.UserToken{}, err
	}

	return result, nil
}
