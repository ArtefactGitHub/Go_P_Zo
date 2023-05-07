package client

import (
	"database/sql"
	"time"
)

type Client interface {
	Id() int
	Secret() string
	CreatedAt() time.Time
	UpdatedAt() sql.NullTime
}

func NewClient(
	id int,
	secret string,
	createdAt time.Time,
	updatedAt sql.NullTime,
) Client {
	return client{
		id:        id,
		secret:    secret,
		createdAt: createdAt,
		updatedAt: updatedAt}
}

type client struct {
	id        int
	secret    string
	createdAt time.Time
	updatedAt sql.NullTime
}

func (c client) Id() int                 { return c.id }
func (c client) Secret() string          { return c.secret }
func (c client) CreatedAt() time.Time    { return c.createdAt }
func (c client) UpdatedAt() sql.NullTime { return c.updatedAt }
