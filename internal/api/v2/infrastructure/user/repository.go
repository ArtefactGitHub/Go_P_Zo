package user

import (
	"context"
	"database/sql"
	"time"

	derr "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
	domain "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/user"
	infra "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure"
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

func (r *repository) FindByIdentifier(ctx context.Context, identifier string, password string) (domain.User, error) {
	db, err := infra.GetDB(ctx)
	if err != nil {
		return domain.User{}, err
	}

	record := userRecord{}
	err = db.QueryRowContext(ctx,
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
		return domain.User{}, derr.NotFound
	} else if err != nil {
		return domain.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(record.Password), []byte(password))
	if err != nil {
		return domain.User{}, err
	}

	return toUser(record), nil
}

func (r *repository) Find(ctx context.Context, id int) (domain.User, error) {
	db, err := infra.GetDB(ctx)
	if err != nil {
		return domain.User{}, err
	}

	record := userRecord{}
	err = db.QueryRowContext(ctx, "SELECT * FROM Users WHERE id = ?", id).Scan(
		&record.Id,
		&record.GivenName,
		&record.FamilyName,
		&record.Email,
		&record.Password,
		&record.CreatedAt,
		&record.UpdatedAt)
	if err == sql.ErrNoRows {
		return domain.User{}, derr.NotFound
	} else if err != nil {
		return domain.User{}, err
	}

	return toUser(record), nil
}

func (r *repository) Create(ctx context.Context, u domain.User) (domain.User, error) {
	password := []byte(u.Password)
	hashed, err := bcrypt.GenerateFromPassword(password, 12)
	if err != nil {
		return domain.User{}, err
	}

	createdAt := time.Now()
	result, err := mydb.Db.ExecContext(ctx, `
			INSERT INTO Users(id, given_name, family_name, email, password, createdAt, updatedAt)
			values(?, ?, ?, ?, ?, ?, ?)`,
		nil,
		&u.GivenName,
		&u.FamilyName,
		&u.Email,
		&hashed,
		&createdAt,
		nil)
	if err != nil {
		return domain.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.User{}, err
	}

	return domain.NewUser(
			int(id),
			u.GivenName,
			u.FamilyName,
			u.Email,
			"",
			createdAt,
			sql.NullTime{}),
		nil
}

func toUser(record userRecord) domain.User {
	result := domain.NewUser(
		record.Id,
		record.GivenName,
		record.FamilyName,
		record.Email,
		record.Password,
		record.CreatedAt,
		record.UpdatedAt,
	)
	return result
}
