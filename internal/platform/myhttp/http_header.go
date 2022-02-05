package myhttp

const (
	UserTokenHeaderName = "X-Go_P_Zo_UserToken"
	AuthTokenHeaderName = "Authorization"
)

type Header struct {
	UserToken string
}
