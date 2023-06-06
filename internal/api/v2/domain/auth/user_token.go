package auth

import (
	"database/sql"
	"time"
)

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
