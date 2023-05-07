package session

import (
	"context"
	"database/sql"

	e "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/session"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	"golang.org/x/crypto/bcrypt"
)

type repository struct {
}

func NewRepository() d.Repository {
	return &repository{}
}

func (r repository) Find(ctx context.Context, identifier string, password string) (d.SessionData, error) {
	result := d.SessionData{}
	storePassword := ""
	err := mydb.Db.QueryRowContext(ctx,
		"SELECT given_name, family_name, Email, Password FROM Users WHERE email = ?", identifier).
		Scan(
			&result.GivenName,
			&result.FamilyName,
			&result.Email,
			&storePassword)
	if err == sql.ErrNoRows {
		return d.SessionData{}, e.NotFound
	} else if err != nil {
		return d.SessionData{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storePassword), []byte(password))
	if err != nil {
		return d.SessionData{}, err
	}

	return result, nil
}
