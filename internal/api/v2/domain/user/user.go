package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type (
	User struct {
		Id         int          `json:"id"`
		GivenName  string       `json:"given_name"`
		FamilyName string       `json:"family_name"`
		Email      string       `json:"email"`
		Password   string       `json:"password"`
		CreatedAt  time.Time    `json:"createdat"`
		UpdatedAt  sql.NullTime `json:"updatedat"`
	}

	Repository interface {
		FindByIdentifier(ctx context.Context, identifier string, password string) (User, error)
		Find(ctx context.Context, id int) (User, error)
	}
)

func NewUser(
	id int,
	givenName string,
	familyName string,
	email string,
	password string,
	createdAt time.Time,
	updatedAt sql.NullTime,
) User {
	return User{
		Id:         id,
		GivenName:  givenName,
		FamilyName: familyName,
		Email:      email,
		Password:   password,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt}
}

func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.FamilyName, u.GivenName)
}
