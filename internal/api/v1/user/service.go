package user

import (
	"context"
	"database/sql"
	"fmt"
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

type UserTokenService struct {
	ur  UserRepository
	utr UserTokenRepository
}

func (s *UserTokenService) Post(ctx context.Context, userId int, m *UserTokenRequest) (*UserToken, error) {
	user, err := s.ur.FindByIdentifier(ctx, m.Identifier, m.Secret)
	if err != nil {
		return nil, err
	}
	if userId != user.Id {
		return nil, fmt.Errorf("invalid userId: %d", userId)
	}

	// TODO
	token := "hoge"
	expiredAt := time.Now().Add(1 * time.Minute)
	data := &UserToken{UserId: user.Id, Token: token, ExpiredAt: expiredAt, CreatedAt: time.Now(), UpdatedAt: sql.NullTime{}}
	_, err = s.utr.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
