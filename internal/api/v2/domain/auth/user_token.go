package auth

import (
	"database/sql"
	"time"
)

type UserToken interface {
	ID() int
	UserID() int
	Token() string
	ExpiredAt() time.Time
	CreatedAt() time.Time
	UpdatedAt() sql.NullTime
}

type userToken struct {
	id        int
	userID    int
	token     string
	expiredAt time.Time
	createdAt time.Time
	updatedAt sql.NullTime
}

func NewUserToken(
	id int,
	userId int,
	token string,
	expiredAt time.Time,
	createdAt time.Time,
	updatedAt sql.NullTime,
) UserToken {
	return userToken{
		id:        id,
		userID:    userId,
		token:     token,
		expiredAt: expiredAt,
		createdAt: createdAt,
		updatedAt: updatedAt}
}

func (v userToken) ID() int {
	return v.id
}
func (v userToken) UserID() int {
	return v.userID
}
func (v userToken) Token() string {
	return v.token
}
func (v userToken) ExpiredAt() time.Time {
	return v.expiredAt
}
func (v userToken) CreatedAt() time.Time {
	return v.createdAt
}
func (v userToken) UpdatedAt() sql.NullTime {
	return v.updatedAt
}
