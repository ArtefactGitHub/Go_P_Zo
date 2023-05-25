package util

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	derr "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mycontext"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

const ResourceIdZo = "zo_id"

func GetUserIdFromToken(ctx context.Context) (int, error) {
	// ユーザートークンの取得
	userToken, err := mycontext.FromContextStr(ctx, mycontext.UserTokenKey)
	if err != nil {
		log.Printf("FromContextStr() err: %#v \n", err)
		return -1, err
	}

	// ユーザートークンからユーザーIDの取得
	claims, err := myauth.CreateUserTokenClaims(userToken)
	if err != nil {
		log.Println(err.Error())
		return -1, err
	}

	return claims.UserId, nil
}

func GetResourceId(params common.QueryMap, resourceKey string) (int, error) {
	idStr, err := GetResourceIdStr(params, resourceKey)
	if err != nil {
		return 0, err
	}

	result, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("リソース指定が正しくありません %s", resourceKey)
	}
	return result, nil
}

func GetResourceIdStr(params common.QueryMap, resourceKey string) (string, error) {
	result := params.Get(resourceKey)
	if result == "" {
		return "", fmt.Errorf("リソース指定が必要です %s", resourceKey)
	}
	return result, nil
}

func HandleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, derr.BadRequest):
		myhttp.WriteError(w, err, http.StatusBadRequest, "リクエストが正しくありません")
	case errors.Is(err, derr.NotFound):
		myhttp.WriteError(w, err, http.StatusNotFound, "リソースが存在しません")
	case errors.Is(err, derr.Unauthorized):
		myhttp.WriteError(w, err, http.StatusUnauthorized, "指定のリソースへアクセスする権限がありません")
	default:
		myhttp.WriteError(w, err, http.StatusInternalServerError, "エラーが発生しました")
	}
}

// Wrap see: https://github.com/golang/pkgsite/blob/master/internal/derrors/derrors.go
func Wrap(err error, format string, args ...interface{}) error {
	result := fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), err)
	return result
}
