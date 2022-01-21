package session

type SessionRequest struct {
	Identifier string `json:"identifier"`
	Secret     string `json:"secret"`
}

func NewSessionRequest(
	identifier string,
	secret string,
) SessionRequest {
	return SessionRequest{
		Identifier: identifier,
		Secret:     secret}
}

type SessionData struct {
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Email      string `json:"email"`
}
