package auth

import (
	"context"
	"database/sql"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/auth"
	du "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/user"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
)

type (
	Create interface {
		Do(context.Context, CreateTokenData) (auth.UserToken, error)
	}
	create struct {
		r  auth.Repository
		ur du.Repository
	}

	CreateTokenData struct {
		Identifier string `json:"identifier"`
		Secret     string `json:"secret"`
	}
)

func NewCreate(r auth.Repository, ur du.Repository) Create {
	return create{r: r, ur: ur}
}

func (u create) Do(ctx context.Context, data CreateTokenData) (auth.UserToken, error) {
	user, err := u.ur.FindByIdentifier(ctx, data.Identifier, data.Secret)
	if err != nil {
		return auth.UserToken{}, err
	}

	ut, err := toCreateUserToken(ctx, user.Id)
	if err != nil {
		return auth.UserToken{}, err
	}

	result, err := u.r.Create(ctx, ut)
	if err != nil {
		return auth.UserToken{}, err
	}

	return result, nil
}

func toCreateUserToken(_ context.Context, userId int) (auth.UserToken, error) {
	expiredAt := time.Now().Add(time.Minute * time.Duration(config.Cfg.Auth.UserTokenExpiration))
	jwt, err := myauth.CreateUserTokenJwt(userId, expiredAt)
	if err != nil {
		return auth.UserToken{}, err
	}

	return auth.NewUserToken(
		0,
		userId,
		jwt,
		expiredAt,
		time.Now(),
		sql.NullTime{},
	), nil
}
