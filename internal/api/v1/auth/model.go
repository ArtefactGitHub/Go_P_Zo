package auth

import (
	"database/sql"
	"time"
)

type Client struct {
	Id        int          `json:"id"`
	Secret    string       `json:"secret"`
	CreatedAt time.Time    `json:"createdat"`
	UpdatedAt sql.NullTime `json:"updatedat"`
}

func NewClient(
	id int,
	secret string,
	createdAt time.Time,
	updatedAt sql.NullTime,
) Client {
	return Client{
		Id:        id,
		Secret:    secret,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt}
}
