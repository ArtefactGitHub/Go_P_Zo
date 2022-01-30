package user

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/middleware"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
)

var (
	route_tests = map[string]func(t *testing.T){
		"test_user_route_getall":    test_user_route_getall,
		"test_user_route_get":       test_user_route_get,
		"test_user_route_post":      test_user_route_post,
		"test_user_route_update":    test_user_route_update,
		"test_user_route_delete":    test_user_route_delete,
		"test_usertoken_route_post": test_usertoken_route_post}

	mockRoutes map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap) = map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap){
		{Path: "/api/v1/users", Method: "GET", NeedAuth: false}:                  uc.getAll,
		{Path: "/api/v1/users/:user_id", Method: "GET", NeedAuth: false}:         uc.get,
		{Path: "/api/v1/users", Method: "POST", NeedAuth: false}:                 uc.post,
		{Path: "/api/v1/users/:user_id", Method: "PUT", NeedAuth: false}:         uc.update,
		{Path: "/api/v1/users/:user_id", Method: "DELETE", NeedAuth: false}:      uc.delete,
		{Path: "/api/v1/users/:user_id/tokens", Method: "POST", NeedAuth: false}: utc.post,
	}

	postUser = NewUser(0, "Bob", "Michael", "bob@gmail.com", "password", time.Now(), sql.NullTime{})
)

func Test_route(t *testing.T) {
	test.Run(t, route_tests, nil, nil, test_seed)
}

// [GET] /api/v1/users のルーティングのテスト
func test_user_route_getall(t *testing.T) {
	writer, err := serveHTTP("GET", "/api/v1/users", nil)
	if err != nil {
		t.Fatalf("serveHTTP failuer. %v", err)
	}

	want := http.StatusOK
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res GetAllResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Fatalf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}

	if len(res.Users) != len(seedUsers) {
		t.Errorf("Invalid Users length. len() = %v, want %d", len(res.Users), len(seedUsers))
	}
}

// [GET] /api/v1/users/:user_id のルーティングのテスト
func test_user_route_get(t *testing.T) {
	writer, err := serveHTTP("GET", "/api/v1/users/1", nil)
	if err != nil {
		t.Fatalf("serveHTTP failuer. %v", err)
	}

	want := http.StatusOK
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res GetResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Fatalf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}

	if res.User.Id != 1 {
		t.Errorf("Invalid User. %v", res.User)
	}
}

// [POST] /api/v1/users のルーティングのテスト
func test_user_route_post(t *testing.T) {
	j, _ := json.MarshalIndent(postUser, "", "\t")

	writer, err := serveHTTP("POST", "/api/v1/users", bytes.NewReader(j))
	if err != nil {
		t.Fatalf("serveHTTP failuer. %v", err)
	}

	want := http.StatusCreated
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res PostResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Fatalf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}

	if res.User.GivenName != postUser.GivenName {
		t.Errorf("Invalid User. %v", res.User)
	}
}

// [UPDATE] /api/v1/users/:user_id のルーティングのテスト
func test_user_route_update(t *testing.T) {
	u := seedUsers[2]
	u.GivenName = "John更新"
	j, _ := json.MarshalIndent(u, "", "\t")

	writer, err := serveHTTP("PUT", "/api/v1/users/1", bytes.NewReader(j))
	if err != nil {
		t.Fatalf("serveHTTP failuer. %v", err)
	}

	want := http.StatusOK
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res PutResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Fatalf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}

	if res.User.GivenName != u.GivenName {
		t.Errorf("Invalid User. %v", res.User)
	}
}

// [DELETE] /api/v1/users/:user_id のルーティングのテスト
func test_user_route_delete(t *testing.T) {
	writer, err := serveHTTP("DELETE", "/api/v1/users/1", nil)
	if err != nil {
		t.Fatalf("serveHTTP failuer. %v", err)
	}

	want := http.StatusOK
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res DeleteResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Errorf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}
}

// [POST] /api/v1/users/:user_id/tokens のルーティングのテスト
func test_usertoken_route_post(t *testing.T) {
	test_user_route_post(t)

	m := NewUserTokenRequest(postUser.Email, postUser.Password)
	j, _ := json.MarshalIndent(m, "", "\t")

	userId := len(seedUsers) + 1
	writer, err := serveHTTP("POST", fmt.Sprintf("/api/v1/users/%d/tokens", userId), bytes.NewReader(j))
	if err != nil {
		t.Fatalf("serveHTTP failuer. %v", err)
	}

	want := http.StatusCreated
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res PostUserTokenResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Fatalf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}

	if res.UserId != len(seedUsers)+1 {
		t.Errorf("Invalid UserToken. %v", res.UserToken)
	}
}

// テスト用のリクエストを実行
func serveHTTP(method string, url string, body io.Reader) (*httptest.ResponseRecorder, error) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest(method, url, body)

	config, _ := test.LoadConfig()
	jwt, err := middleware.NewJwtMiddleware(config)
	if err != nil {
		return nil, err
	}

	handler := middleware.CreateHandler(
		jwt,
		middleware.NewRouterMiddleware(
			mockRoutes,
		),
	)

	handler.ServeHTTP(writer, request)
	return writer, nil
}
