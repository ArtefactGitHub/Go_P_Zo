package user

import (
	"database/sql"
	"time"
)

type User struct {
	Id         int          `json:"id"`
	GivenName  string       `json:"given_name"`
	FamilyName string       `json:"family_name"`
	Email      string       `json:"email"`
	CreatedAt  time.Time    `json:"createdat"`
	UpdatedAt  sql.NullTime `json:"updatedat"`
}

func NewUser(
	id int,
	givenName string,
	familyName string,
	email string,
	createdAt time.Time,
	updatedAt sql.NullTime,
) User {
	return User{
		Id:         id,
		GivenName:  givenName,
		FamilyName: familyName,
		Email:      email,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt}
}
