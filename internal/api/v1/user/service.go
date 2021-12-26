package user

import (
	"context"
)

type UserService struct {
	Ur UserRepository
}

func (s *UserService) GetAll(ctx context.Context) ([]User, error) {
	result, err := s.Ur.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UserService) Get(ctx context.Context, id int) (*User, error) {
	result, err := s.Ur.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UserService) Post(ctx context.Context, z *User) (int, error) {
	result, err := s.Ur.Create(ctx, z)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (s *UserService) Update(ctx context.Context, z *User) error {
	err := s.Ur.Update(ctx, z)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Delete(ctx context.Context, id int) error {
	err := s.Ur.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
