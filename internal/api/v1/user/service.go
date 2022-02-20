package user

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
)

// User
type UserService struct {
	r UserRepository
}

func (s *UserService) GetAll(ctx context.Context) ([]responseUser, error) {
	users, err := s.r.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	result := NewResponseUsers(users)
	return result, nil
}

func (s *UserService) Get(ctx context.Context, id int) (*responseUser, error) {
	user, err := s.r.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	result := NewResponseUser(user.Id, user.GivenName, user.FamilyName, user.Email)
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

// UserToken
type userTokenService struct {
	ur  UserRepository
	utr UserTokenRepository
}

func (s *userTokenService) Post(ctx context.Context, m *UserTokenRequest) (*UserToken, error) {
	user, err := s.ur.FindByIdentifier(ctx, m.Identifier, m.Secret)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("認証情報が正しくありません。")
	}

	userToken, err := s.createUserToken(user.Id)
	if err != nil {
		return nil, err
	}

	_, err = s.utr.Create(ctx, userToken)
	if err != nil {
		return nil, err
	}

	return userToken, nil
}

func (s *userTokenService) createUserToken(userId int) (*UserToken, error) {
	expiredAt := time.Now().Add(time.Duration(time.Minute * time.Duration(config.Cfg.Auth.UserTokenExpiration)))
	jwt, err := myauth.CreateUserTokenJwt(userId, expiredAt)
	if err != nil {
		return nil, err
	}

	result := &UserToken{UserId: userId, Token: jwt, ExpiredAt: expiredAt, CreatedAt: time.Now(), UpdatedAt: sql.NullTime{}}
	return result, nil
}

// userCategory
type userCategoryService struct {
	r userCategoryRepository
}

func (s *userCategoryService) GetAll(ctx context.Context, userId int) ([]responseUserCategory, error) {
	models, err := s.r.FindAllByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	result := NewResponseUserCategories(models)
	return result, nil
}

func (s *userCategoryService) Post(ctx context.Context, userId int, r *requestUserCategory) (int, error) {
	m := NewUserCategory(0, 0, r.Name, r.ColorId, userId, time.Now(), sql.NullTime{})
	result, err := s.r.Create(ctx, m)
	if err != nil {
		return -1, err
	}
	return result, nil
}
