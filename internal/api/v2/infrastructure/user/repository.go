package user

import (
	"context"
	"database/sql"
	"time"

	derr "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
	domain "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/user"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	"golang.org/x/crypto/bcrypt"
)

type (
	repository struct {
	}
	userRecord struct {
		Id         int          `json:"id"`
		GivenName  string       `json:"given_name"`
		FamilyName string       `json:"family_name"`
		Email      string       `json:"email"`
		Password   string       `json:"password"`
		CreatedAt  time.Time    `json:"createdat"`
		UpdatedAt  sql.NullTime `json:"updatedat"`
	}
)

func NewRepository() domain.Repository {
	return &repository{}
}

func (r *repository) FindByIdentifier(ctx context.Context, identifier string, password string) (*domain.User, error) {
	record := userRecord{}
	err := mydb.Db.QueryRowContext(ctx,
		"SELECT * FROM Users WHERE email = ?", identifier).
		Scan(
			&record.Id,
			&record.GivenName,
			&record.FamilyName,
			&record.Email,
			&record.Password,
			&record.CreatedAt,
			&record.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, derr.NotFound
	} else if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(record.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return toUser(record), nil
}

func toUser(record userRecord) *domain.User {
	result := domain.NewUser(
		record.Id,
		record.GivenName,
		record.FamilyName,
		record.Email,
		record.Password,
		record.CreatedAt,
		record.UpdatedAt,
	)
	return &result
}
