package client

import (
	"database/sql"
	"time"
)

const (
	Issuer    = "zo.auth.service"
	TokenType = "accessToken"
)

type AccessToken interface {
	Jwt() string
	ExpiresAt() int64
}

type accessToken struct {
	jwt       string `json:"jwt"`
	expiresAt int64  `json:"expires_at"`
}

func NewAccessToken(jwt string, expiresAt int64) AccessToken {
	return accessToken{
		jwt:       jwt,
		expiresAt: expiresAt,
	}
}

func (t accessToken) Jwt() string      { return t.jwt }
func (t accessToken) ExpiresAt() int64 { return t.expiresAt }

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
	id        int          `json:"id"`
	secret    string       `json:"secret"`
	createdAt time.Time    `json:"create_date"`
	updatedAt sql.NullTime `json:"update_date"`
}

func (c client) Id() int                 { return c.id }
func (c client) Secret() string          { return c.secret }
func (c client) CreatedAt() time.Time    { return c.createdAt }
func (c client) UpdatedAt() sql.NullTime { return c.updatedAt }
