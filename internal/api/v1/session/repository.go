package session

import (
	"context"
	"database/sql"
	"log"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	"golang.org/x/crypto/bcrypt"
)

type SessionRepository struct {
}

func (r *SessionRepository) Find(ctx context.Context, identifier string, password string) (*SessionData, error) {
	result := &SessionData{}
	storePassword := ""
	err := mydb.Db.QueryRowContext(ctx,
		"SELECT given_name, family_name, Email, Password FROM users WHERE email = ?", identifier).
		Scan(
			&result.GivenName,
			&result.FamilyName,
			&result.Email,
			&storePassword)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storePassword), []byte(password))
	if err != nil {
		return nil, err
	}
	log.Printf("fff  %v", result)
	return result, nil
}
