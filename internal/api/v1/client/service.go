package client

import (
	"context"
)

type clientService struct {
	r clientRepository
}

func (s *clientService) Exist(ctx context.Context, id int, secret string) (bool, error) {
	result, err := s.r.Find(ctx, id, secret)
	if err != nil {
		return false, err
	}

	return result != nil, nil
}

func (s *clientService) Find(ctx context.Context, id int, secret string) (*Client, error) {
	result, err := s.r.Find(ctx, id, secret)
	if err != nil {
		return nil, err
	}

	return result, nil
}
