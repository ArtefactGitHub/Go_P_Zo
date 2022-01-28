package user

import (
	"context"
	"database/sql"
	"time"
)

type UserService struct {
	r UserRepository
}

func (s *UserService) GetAll(ctx context.Context) ([]User, error) {
	result, err := s.r.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UserService) Get(ctx context.Context, id int) (*User, error) {
	result, err := s.r.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UserService) Post(ctx context.Context, u *User) (int, error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = sql.NullTime{}
	result, err := s.r.Create(ctx, u)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (s *UserService) Update(ctx context.Context, u *User) error {
	u.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	err := s.r.Update(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Delete(ctx context.Context, id int) error {
	err := s.r.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
