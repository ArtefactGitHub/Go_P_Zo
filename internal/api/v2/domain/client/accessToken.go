package client

type AccessToken interface {
	Jwt() string
	ExpiresAt() int64
}

type accessToken struct {
	jwt       string
	expiresAt int64
}

func NewAccessToken(jwt string, expiresAt int64) AccessToken {
	return accessToken{
		jwt:       jwt,
		expiresAt: expiresAt,
	}
}

func (t accessToken) Jwt() string      { return t.jwt }
func (t accessToken) ExpiresAt() int64 { return t.expiresAt }
