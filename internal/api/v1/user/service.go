package user

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
	"github.com/golang-jwt/jwt"
)

// User
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
	expiredAt := time.Now().Add(time.Minute * time.Duration(config.Cfg.Auth.UserTokenExpiration))
	claims := myauth.UserTokenClaims{StandardClaims: &jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Issuer:    "zo.auth.service",
	},
		TokenType: "userToken",
		UserId:    userId,
	}

	// https://pkg.go.dev/github.com/golang-jwt/jwt#NewWithClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := token.SignedString([]byte(config.Cfg.Auth.SignKey))
	log.Printf("signed usertoken: %v", jwt)

	if err != nil {
		return nil, err
	}

	result := &UserToken{UserId: userId, Token: jwt, ExpiredAt: expiredAt, CreatedAt: time.Now(), UpdatedAt: sql.NullTime{}}
	return result, nil
}
