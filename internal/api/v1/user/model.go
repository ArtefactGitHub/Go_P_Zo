package user

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	Id         int          `json:"id"`
	GivenName  string       `json:"given_name"`
	FamilyName string       `json:"family_name"`
	Email      string       `json:"email"`
	Password   string       `json:"password"`
	CreatedAt  time.Time    `json:"createdat"`
	UpdatedAt  sql.NullTime `json:"updatedat"`
}

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

type UserToken struct {
	Id        int          `json:"id"`
	UserId    int          `json:"user_id"`
	Token     string       `json:"token"`
	ExpiredAt time.Time    `json:"expiredat"`
	CreatedAt time.Time    `json:"createdat"`
	UpdatedAt sql.NullTime `json:"updatedat"`
}

func NewUserToken(
	id int,
	userId int,
	token string,
	expiredAt time.Time,
	createdAt time.Time,
	updatedAt sql.NullTime,
) UserToken {
	return UserToken{
		Id:        id,
		UserId:    userId,
		Token:     token,
		ExpiredAt: expiredAt,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt}
}

type UserCategory struct {
	Id        int          `json:"id"`
	Number    int          `json:"number"`
	Name      string       `json:"name"`
	ColorId   int          `json:"color_id"`
	UserId    int          `json:"user_id"`
	CreatedAt time.Time    `json:"createdat"`
	UpdatedAt sql.NullTime `json:"updatedat"`
}

func NewUserCategory(
	id int,
	number int,
	name string,
	colorId int,
	userId int,
	createdAt time.Time,
	updatedAt sql.NullTime,
) *UserCategory {
	return &UserCategory{
		Id:        id,
		Number:    number,
		Name:      name,
		ColorId:   colorId,
		UserId:    userId,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt}
}
