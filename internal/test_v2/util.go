package test_v2

import (
	"context"
	"database/sql"
	"path/filepath"
	"runtime"
	"time"

	infra "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure"
	util "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/presentation/utils"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mycontext"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

type TrashScanner struct{}

func (TrashScanner) Scan(interface{}) error {
	return nil
}

func currentDir() string {
	_, pwd, _, _ := runtime.Caller(0)
	return filepath.Dir(pwd)
}

func GetResourceId(params common.QueryMap, resourceKey string) int {
	v, _ := util.GetResourceId(params, resourceKey)
	return v
}

func GetResourceIdStr(params common.QueryMap, resourceKey string) string {
	v, _ := util.GetResourceIdStr(params, resourceKey)
	return v
}

func WithDBAndTokenContext(ctx context.Context, db *sql.DB, userID int, expiration time.Time) context.Context {
	c := WithDBContext(ctx, db)
	return WithTokenContext(c, userID, expiration)
}

func WithDBContext(ctx context.Context, db *sql.DB) context.Context {
	return context.WithValue(ctx, infra.KeyDB, db)
}

func WithTokenContext(ctx context.Context, userID int, expiration time.Time) context.Context {
	jwt, _ := myauth.CreateUserTokenJwt(userID, expiration)
	return mycontext.NewContext(ctx, mycontext.UserTokenKey, jwt)
}
