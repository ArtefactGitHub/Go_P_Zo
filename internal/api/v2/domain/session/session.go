package session

type SessionData struct {
	GivenName  string
	FamilyName string
	Email      string
}

func NewSessionData(givenName, familyName, email string) SessionData {
	return SessionData{
		GivenName:  givenName,
		FamilyName: familyName,
		Email:      email,
	}
}
