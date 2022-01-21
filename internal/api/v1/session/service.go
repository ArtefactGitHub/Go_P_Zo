package session

import (
	"context"
)

type SessionService struct {
	r SessionRepository
}

func (s *SessionService) Login(ctx context.Context, identifier string, secret string) (*SessionData, error) {
	result, err := s.r.Find(ctx, identifier, secret)
	if err != nil {
		return nil, err
	}

	return result, nil
}
