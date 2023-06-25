package auth

import (
	"context"
	"database/sql"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/auth"
	e "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
	infra "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure"
)

type (
	repository struct{}

	userTokenRecord struct {
		ID        int          `json:"id"`
		UserID    int          `json:"user_id"`
		Token     string       `json:"token"`
		ExpiredAt time.Time    `json:"expiredat"`
		CreatedAt time.Time    `json:"createdat"`
		UpdatedAt sql.NullTime `json:"updatedat"`
	}
)

func NewRepository() auth.Repository {
	return &repository{}
}

func (r *repository) Create(ctx context.Context, v auth.UserToken) (auth.UserToken, error) {
	tx, err := infra.GetTX(ctx)
	if err != nil {
		return nil, err
	}

	record := toRecord(v)
	result, err := tx.ExecContext(ctx, `
		INSERT INTO UserTokens(id, user_id, token, expiredAt, createdAt, updatedAt)
			VALUES(?,?,?,?,?,?)`,
		nil,
		&record.UserID,
		&record.Token,
		&record.ExpiredAt,
		&record.CreatedAt,
		&record.UpdatedAt)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return auth.NewUserToken(
			int(id),
			record.UserID,
			record.Token,
			record.ExpiredAt,
			record.CreatedAt,
			record.UpdatedAt),
		nil
}

func (r repository) Find(ctx context.Context, id int) (auth.UserToken, error) {
	db, err := infra.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	record := userTokenRecord{}
	err = db.QueryRowContext(ctx, "SELECT * FROM UserTokens WHERE id = ?",
		id).
		Scan(
			&record.ID,
			&record.UserID,
			&record.Token,
			&record.ExpiredAt,
			&record.CreatedAt,
			&record.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, e.NotFound
	} else if err != nil {
		return nil, err
	}

	return auth.NewUserToken(
			id,
			record.UserID,
			record.Token,
			record.ExpiredAt,
			record.CreatedAt,
			record.UpdatedAt),
		nil
}

func toRecord(v auth.UserToken) userTokenRecord {
	return userTokenRecord{
		ID:        v.ID(),
		UserID:    v.UserID(),
		Token:     v.Token(),
		ExpiredAt: v.ExpiredAt(),
		CreatedAt: v.CreatedAt(),
		UpdatedAt: v.UpdatedAt(),
	}
}
