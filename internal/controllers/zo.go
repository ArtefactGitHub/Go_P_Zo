package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/models"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/models/zo"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/services"
)

var resourceName string = "zo"

type ZoController struct {
	zs services.ZoService
}

func (zc *ZoController) ZoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r)

	switch r.Method {
	case http.MethodGet:
		zc.get(w, r)
	case http.MethodPost:
		zc.post(w, r)
	case http.MethodPut, http.MethodPatch:
		zc.update(w, r)
	case http.MethodDelete:
		zc.delete(w, r)
	default:
		WriteError(w, nil, http.StatusMethodNotAllowed, fmt.Sprintf("%s method not allowed", r.Method))
	}
}

// リソースを取得
func (c *ZoController) get(w http.ResponseWriter, r *http.Request) {
	// リソース群の取得
	if path.Base(r.URL.Path) == resourceName {
		datas, err := c.zs.GetAll()
		if err != nil {
			WriteError(w, err, http.StatusInternalServerError, "")
			return
		}

		res := zo.GetAllResponse{
			ResponseBase: models.ResponseBase{StatusCode: http.StatusOK, Error: nil},
			Zos:          datas}
		WriteSuccess(w, res, http.StatusOK)
	} else {
		// 指定リソースの取得
		// 末尾のid指定を取得
		id, err := strconv.Atoi(path.Base(r.URL.Path))
		if err != nil {
			WriteError(w, err, http.StatusBadRequest, "incorrect resource specification")
			return
		}

		model, err := c.zs.Get(id)
		if err != nil {
			WriteError(w, err, http.StatusInternalServerError, "")
			return
		} else if model == nil {
			WriteError(w, fmt.Errorf("resource not found: id = %d", id),
				http.StatusNotFound, "")
			return
		}

		res := zo.GetResponse{
			ResponseBase: models.ResponseBase{StatusCode: http.StatusOK, Error: nil},
			Zo:           model}
		WriteSuccess(w, res, http.StatusOK)
	}
}

// 指定のリソース情報で作成
func (c *ZoController) post(w http.ResponseWriter, r *http.Request) {
	// リクエスト情報からモデルを生成
	m, err := contentToModel(r)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	id, err := c.zs.Post(m)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := zo.PostResponse{
		ResponseBase: models.ResponseBase{StatusCode: http.StatusOK, Error: nil},
		Zo:           m}
	WriteSuccessWithLocation(w, res, http.StatusCreated, r.Host+r.URL.Path+strconv.Itoa(id))
}

// 指定のリソース情報で更新
func (c *ZoController) update(w http.ResponseWriter, r *http.Request) {
	// 末尾のid指定を取得
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	// リクエスト情報からモデルを生成
	m, err := contentToModel(r)
	log.Printf("contentToModel: %v", m)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	u := zo.NewZo(
		id,
		m.AchievementDate,
		m.Exp,
		m.CategoryId,
		m.Message,
		m.CreatedAt,
		sql.NullTime{Time: time.Now(), Valid: true})
	u, err = c.zs.Update(u)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := zo.PutResponse{
		ResponseBase: models.ResponseBase{StatusCode: http.StatusOK, Error: nil},
		Zo:           u}
	WriteSuccess(w, res, http.StatusOK)
}

// 指定のリソースの削除
func (c *ZoController) delete(w http.ResponseWriter, r *http.Request) {
	// 指定リソースの取得
	// 末尾のid指定を取得
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		WriteError(w, nil, http.StatusBadRequest, "incorrect resource specification")
		return
	}

	err = c.zs.Delete(id)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := zo.DeleteResponse{ResponseBase: models.ResponseBase{StatusCode: http.StatusOK, Error: nil}}
	WriteSuccess(w, res, http.StatusOK)
}

// リクエスト情報からモデルの生成
func contentToModel(r *http.Request) (*zo.Zo, error) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var result zo.Zo
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
