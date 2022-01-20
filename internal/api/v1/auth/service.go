package auth

import (
	"context"
)

type AuthService struct {
	r AuthRepository
}

func (s *AuthService) Exist(ctx context.Context, id int, secret string) (bool, error) {
	result, err := s.r.Find(ctx, id, secret)
	if err != nil {
		return false, err
	}

	return result != nil, nil
}

func (s *AuthService) Find(ctx context.Context, id int, secret string) (*Client, error) {
	result, err := s.r.Find(ctx, id, secret)
	if err != nil {
		return nil, err
	}

	return result, nil
}
