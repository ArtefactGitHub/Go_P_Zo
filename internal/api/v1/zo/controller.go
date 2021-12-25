package zo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
	"github.com/julienschmidt/httprouter"
)

type zoController struct {
	zs ZoService
}

const resourceId = "zo_id"

// リソースを取得
func (c *zoController) getAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// リソース群の取得
	datas, err := c.zs.GetAll(r.Context())
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := GetAllResponse{
		ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusOK, Error: nil},
		Zos:          datas}
	myhttp.WriteSuccess(w, res, http.StatusOK)
}

func (c *zoController) get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 指定リソースの取得
	// 末尾のid指定を取得
	id, err := strconv.Atoi(ps.ByName(resourceId))
	if err != nil {
		myhttp.WriteError(w, err, http.StatusBadRequest, "incorrect resource specification")
		return
	}

	model, err := c.zs.Get(r.Context(), id)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	} else if model == nil {
		myhttp.WriteError(w, fmt.Errorf("resource not found: id = %d", id),
			http.StatusNotFound, "")
		return
	}

	res := GetResponse{
		ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusOK, Error: nil},
		Zo:           model}
	myhttp.WriteSuccess(w, res, http.StatusOK)
}

// 指定のリソース情報で作成
func (c *zoController) post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// リクエスト情報からモデルを生成
	m, err := contentToModel(r)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	id, err := c.zs.Post(r.Context(), m)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := PostResponse{
		ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusCreated, Error: nil},
		Zo:           m}
	myhttp.WriteSuccessWithLocation(w, res, http.StatusCreated, r.Host+r.URL.Path+strconv.Itoa(id))
}

// 指定のリソース情報で更新
func (c *zoController) update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 末尾のid指定を取得
	id, err := strconv.Atoi(ps.ByName(resourceId))
	if err != nil {
		myhttp.WriteError(w, err, http.StatusBadRequest, "incorrect resource specification")
		return
	}

	// リクエスト情報からモデルを生成
	m, err := contentToModel(r)
	log.Printf("contentToModel: %v", m)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	u := NewZo(
		id,
		m.AchievementDate,
		m.Exp,
		m.CategoryId,
		m.Message,
		m.CreatedAt,
		sql.NullTime{Time: time.Now(), Valid: true},
		m.UserId)
	err = c.zs.Update(r.Context(), &u)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := PutResponse{
		ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusOK, Error: nil},
		Zo:           &u}
	myhttp.WriteSuccess(w, res, http.StatusOK)
}

// 指定のリソースの削除
func (c *zoController) delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 指定リソースの取得
	// 末尾のid指定を取得
	id, err := strconv.Atoi(ps.ByName(resourceId))
	if err != nil {
		myhttp.WriteError(w, err, http.StatusBadRequest, "incorrect resource specification")
		return
	}

	err = c.zs.Delete(r.Context(), id)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := DeleteResponse{ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusOK, Error: nil}}
	myhttp.WriteSuccess(w, res, http.StatusOK)
}

// リクエスト情報からモデルの生成
func contentToModel(r *http.Request) (*Zo, error) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var result Zo
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
