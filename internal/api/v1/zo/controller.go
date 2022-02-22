package zo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mycontext"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

type zoController struct {
	zs ZoService
}

const resourceId = "zo_id"

// リソースを取得
func (c *zoController) getAll(w http.ResponseWriter, r *http.Request, params common.QueryMap) {
	// ユーザーIDの取得
	userId, err := c.getUserIdFromToken(r.Context())
	if err != nil {
		myhttp.WriteError(w, err, http.StatusUnauthorized, "指定のリソースへアクセスする権限がありません")
		return
	}

	// リソース群の取得
	datas, err := c.zs.GetAll(r.Context(), userId)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := GetAllResponse{
		ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusOK, Error: nil},
		Zos:          datas}
	myhttp.Write(w, res, http.StatusOK)
}

func (c *zoController) get(w http.ResponseWriter, r *http.Request, params common.QueryMap) {
	// ユーザーIDの取得
	userId, err := c.getUserIdFromToken(r.Context())
	if err != nil {
		myhttp.WriteError(w, err, http.StatusUnauthorized, "指定のリソースへアクセスする権限がありません")
		return
	}

	// 指定リソースの取得
	// 末尾のid指定を取得
	id, err := strconv.Atoi(params.Get(resourceId))
	if err != nil {
		myhttp.WriteError(w, err, http.StatusBadRequest, "リソース指定が正しくありません")
		return
	}

	m, err := c.zs.Get(r.Context(), id)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	} else if m == nil {
		myhttp.WriteError(w, fmt.Errorf("リソースが見つかりません: id = %d", id), http.StatusNotFound, "")
		return
	}

	// 非リソース所有者の場合
	if userId != m.UserId {
		myhttp.WriteError(w, err, http.StatusUnauthorized, "指定のリソースへアクセスする権限がありません")
		return
	}

	res := GetResponse{
		ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusOK, Error: nil},
		Zo:           m}
	myhttp.Write(w, res, http.StatusOK)
}

// 指定のリソース情報で作成
func (c *zoController) post(w http.ResponseWriter, r *http.Request, params common.QueryMap) {
	// ユーザーIDの取得
	userId, err := c.getUserIdFromToken(r.Context())
	if err != nil {
		myhttp.WriteError(w, err, http.StatusUnauthorized, "指定のリソースへアクセスする権限がありません")
		return
	}

	// リクエスト情報からモデルを生成
	m := &requestZo{}
	err = contentToModel(r, m)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	id, err := c.zs.Post(r.Context(), userId, m)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := PostResponse{
		ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusCreated, Error: nil},
		Zo:           NewResponseZo(id, m.AchievementDate, m.Exp, m.CategoryId, m.Message)}
	myhttp.WriteSuccessWithLocation(w, res, http.StatusCreated, r.Host+r.URL.Path+strconv.Itoa(id))
}

// 指定のリソース情報で更新
func (c *zoController) update(w http.ResponseWriter, r *http.Request, params common.QueryMap) {
	// ユーザーIDの取得
	userId, err := c.getUserIdFromToken(r.Context())
	if err != nil {
		myhttp.WriteError(w, err, http.StatusUnauthorized, "指定のリソースへアクセスする権限がありません")
		return
	}

	// 末尾のid指定を取得
	id, err := strconv.Atoi(params.Get(resourceId))
	if err != nil {
		myhttp.WriteError(w, err, http.StatusBadRequest, "リソース指定が正しくありません")
		return
	}

	// リクエスト情報からモデルを生成
	m := &Zo{}
	err = contentToModel(r, m)
	log.Printf("contentToModel: %v", m)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	// 非リソース所有者の場合
	if userId != m.UserId {
		myhttp.WriteError(w, err, http.StatusUnauthorized, "指定のリソースへアクセスする権限がありません")
		return
	}

	z := NewZo(
		id,
		m.AchievementDate,
		m.Exp,
		m.CategoryId,
		m.Message,
		m.CreatedAt,
		sql.NullTime{Time: time.Now(), Valid: true},
		m.UserId)
	err = c.zs.Update(r.Context(), &z)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := PutResponse{
		ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusOK, Error: nil},
		Zo:           &z}
	myhttp.Write(w, res, http.StatusOK)
}

// 指定のリソースの削除
func (c *zoController) delete(w http.ResponseWriter, r *http.Request, params common.QueryMap) {
	// ユーザーIDの取得
	userId, err := c.getUserIdFromToken(r.Context())
	if err != nil {
		myhttp.WriteError(w, err, http.StatusUnauthorized, "指定のリソースへアクセスする権限がありません")
		return
	}

	// 指定リソースの取得
	// 末尾のid指定を取得
	id, err := strconv.Atoi(params.Get(resourceId))
	if err != nil {
		myhttp.WriteError(w, err, http.StatusBadRequest, "incorrect resource specification")
		return
	}

	// リソース所有者かチェック
	m, err := c.zs.Get(r.Context(), id)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	} else if m == nil {
		myhttp.WriteError(w, fmt.Errorf("リソースが見つかりません: id = %d", id), http.StatusNotFound, "")
		return
	}
	if userId != m.UserId {
		myhttp.WriteError(w, err, http.StatusUnauthorized, "指定のリソースへアクセスする権限がありません")
		return
	}

	err = c.zs.Delete(r.Context(), id)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := DeleteResponse{ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusOK, Error: nil}}
	myhttp.Write(w, res, http.StatusOK)
}

func (c *zoController) getUserIdFromToken(ctx context.Context) (int, error) {
	// ユーザートークンの取得
	userToken, err := mycontext.FromContextStr(ctx, mycontext.UserTokenKey)
	if err != nil {
		return -1, err
	}
	log.Print(userToken)

	// ユーザートークンからユーザーIDの取得
	claims, err := myauth.CreateUserTokenClaims(userToken)
	if err != nil {
		log.Println(err.Error())
		return -1, err
	}

	return claims.UserId, nil
}

// リクエスト情報からモデルの生成
func contentToModel(r *http.Request, model interface{}) error {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	err := json.Unmarshal(body, model)
	if err != nil {
		return err
	}
	return nil
}
