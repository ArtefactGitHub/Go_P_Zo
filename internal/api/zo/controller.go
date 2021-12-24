package zo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
)

const resourceName string = "zo"

type zoController struct {
	zs zoService
}

func (c *zoController) handle(w http.ResponseWriter, r *http.Request) {
	log.Println(r)

	switch r.Method {
	case http.MethodGet:
		c.get(w, r)
	case http.MethodPost:
		c.post(w, r)
	case http.MethodPut, http.MethodPatch:
		c.update(w, r)
	case http.MethodDelete:
		c.delete(w, r)
	default:
		myhttp.WriteError(w, nil, http.StatusMethodNotAllowed, fmt.Sprintf("%s method not allowed", r.Method))
	}
}

// リソースを取得
func (c *zoController) get(w http.ResponseWriter, r *http.Request) {
	// リソース群の取得
	if path.Base(r.URL.Path) == resourceName {
		datas, err := c.zs.getAll()
		if err != nil {
			myhttp.WriteError(w, err, http.StatusInternalServerError, "")
			return
		}

		res := getAllResponse{
			ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusOK, Error: nil},
			Zos:          datas}
		myhttp.WriteSuccess(w, res, http.StatusOK)
	} else {
		// 指定リソースの取得
		// 末尾のid指定を取得
		id, err := strconv.Atoi(path.Base(r.URL.Path))
		if err != nil {
			myhttp.WriteError(w, err, http.StatusBadRequest, "incorrect resource specification")
			return
		}

		model, err := c.zs.get(id)
		if err != nil {
			myhttp.WriteError(w, err, http.StatusInternalServerError, "")
			return
		} else if model == nil {
			myhttp.WriteError(w, fmt.Errorf("resource not found: id = %d", id),
				http.StatusNotFound, "")
			return
		}

		res := getResponse{
			ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusOK, Error: nil},
			Zo:           model}
		myhttp.WriteSuccess(w, res, http.StatusOK)
	}
}

// 指定のリソース情報で作成
func (c *zoController) post(w http.ResponseWriter, r *http.Request) {
	// リクエスト情報からモデルを生成
	m, err := contentToModel(r)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	id, err := c.zs.post(m)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := postResponse{
		ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusCreated, Error: nil},
		Zo:           m}
	myhttp.WriteSuccessWithLocation(w, res, http.StatusCreated, r.Host+r.URL.Path+strconv.Itoa(id))
}

// 指定のリソース情報で更新
func (c *zoController) update(w http.ResponseWriter, r *http.Request) {
	// 末尾のid指定を取得
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	// リクエスト情報からモデルを生成
	m, err := contentToModel(r)
	log.Printf("contentToModel: %v", m)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	u := newZo(
		id,
		m.AchievementDate,
		m.Exp,
		m.CategoryId,
		m.Message,
		m.CreatedAt,
		sql.NullTime{Time: time.Now(), Valid: true})
	err = c.zs.update(&u)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := putResponse{
		ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusOK, Error: nil},
		Zo:           &u}
	myhttp.WriteSuccess(w, res, http.StatusOK)
}

// 指定のリソースの削除
func (c *zoController) delete(w http.ResponseWriter, r *http.Request) {
	// 指定リソースの取得
	// 末尾のid指定を取得
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		myhttp.WriteError(w, nil, http.StatusBadRequest, "incorrect resource specification")
		return
	}

	err = c.zs.delete(id)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := deleteResponse{ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusOK, Error: nil}}
	myhttp.WriteSuccess(w, res, http.StatusOK)
}

// リクエスト情報からモデルの生成
func contentToModel(r *http.Request) (*zo, error) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var result zo
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
