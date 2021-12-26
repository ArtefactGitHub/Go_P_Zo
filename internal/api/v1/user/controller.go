package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myhttp"
	"github.com/julienschmidt/httprouter"
)

type userController struct {
	s UserService
}

const resourceId = "user_id"

// リソースを取得
func (c *userController) getAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// リソース群の取得
	datas, err := c.s.GetAll(r.Context())
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := GetAllResponse{
		ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusOK, Error: nil},
		Users:        datas}
	myhttp.WriteSuccess(w, res, http.StatusOK)
}

func (c *userController) get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 指定リソースの取得
	// 末尾のid指定を取得
	id, err := strconv.Atoi(ps.ByName(resourceId))
	if err != nil {
		myhttp.WriteError(w, err, http.StatusBadRequest, "incorrect resource specification")
		return
	}

	model, err := c.s.Get(r.Context(), id)
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
		User:         model}
	myhttp.WriteSuccess(w, res, http.StatusOK)
}

// 指定のリソース情報で作成
func (c *userController) post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// リクエスト情報からモデルを生成
	u, err := contentToModel(r)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	id, err := c.s.Post(r.Context(), u)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := PostResponse{
		ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusCreated, Error: nil},
		User:         u}
	myhttp.WriteSuccessWithLocation(w, res, http.StatusCreated, r.Host+r.URL.Path+strconv.Itoa(id))
}

// 指定のリソース情報で更新
func (c *userController) update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 末尾のid指定を取得
	_, err := strconv.Atoi(ps.ByName(resourceId))
	if err != nil {
		myhttp.WriteError(w, err, http.StatusBadRequest, "incorrect resource specification")
		return
	}

	// リクエスト情報からモデルを生成
	u, err := contentToModel(r)
	log.Printf("contentToModel: %v", u)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	err = c.s.Update(r.Context(), u)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := PutResponse{
		ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusOK, Error: nil},
		User:         u}
	myhttp.WriteSuccess(w, res, http.StatusOK)
}

// 指定のリソースの削除
func (c *userController) delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 指定リソースの取得
	// 末尾のid指定を取得
	id, err := strconv.Atoi(ps.ByName(resourceId))
	if err != nil {
		myhttp.WriteError(w, err, http.StatusBadRequest, "incorrect resource specification")
		return
	}

	err = c.s.Delete(r.Context(), id)
	if err != nil {
		myhttp.WriteError(w, err, http.StatusInternalServerError, "")
		return
	}

	res := DeleteResponse{ResponseBase: myhttp.ResponseBase{StatusCode: http.StatusOK, Error: nil}}
	myhttp.WriteSuccess(w, res, http.StatusOK)
}

// リクエスト情報からモデルの生成
func contentToModel(r *http.Request) (*User, error) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var result User
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
