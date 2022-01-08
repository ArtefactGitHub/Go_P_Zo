package user_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/user"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
)

var route_tests = map[string]func(t *testing.T){
	"test_user_route_getall": test_user_route_getall,
	"test_user_route_get":    test_user_route_get,
	"test_user_route_post":   test_user_route_post,
	"test_user_route_update": test_user_route_update,
	"test_user_route_delete": test_user_route_delete}

func Test_route(t *testing.T) {
	test.Run(t, route_tests, nil, nil, test_seed)
}

// [GET] /api/v1/users のルーティングのテスト
func test_user_route_getall(t *testing.T) {
	writer := serveHTTP("GET", "/api/v1/users", nil)

	want := http.StatusOK
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res user.GetAllResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Fatalf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}

	if len(res.Users) != len(seeds) {
		t.Errorf("Invalid Users length. len() = %v, want %d", len(res.Users), len(seeds))
	}
}

// [GET] /api/v1/users/:user_id のルーティングのテスト
func test_user_route_get(t *testing.T) {
	writer := serveHTTP("GET", "/api/v1/users/1", nil)

	want := http.StatusOK
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res user.GetResponse
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
	u := user.NewUser(1, "Bob", "Michael", "bob@gmail.com", time.Now(), sql.NullTime{})
	j, _ := json.MarshalIndent(u, "", "\t")

	writer := serveHTTP("POST", "/api/v1/users", bytes.NewReader(j))

	want := http.StatusCreated
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res user.PostResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Fatalf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}

	if res.User.GivenName != u.GivenName {
		t.Errorf("Invalid User. %v", res.User)
	}
}

// [UPDATE] /api/v1/users/:user_id のルーティングのテスト
func test_user_route_update(t *testing.T) {
	u := seeds[2]
	u.GivenName = "John更新"
	j, _ := json.MarshalIndent(u, "", "\t")

	writer := serveHTTP("PUT", "/api/v1/users/1", bytes.NewReader(j))

	want := http.StatusOK
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res user.PutResponse
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
	writer := serveHTTP("DELETE", "/api/v1/users/1", nil)

	want := http.StatusOK
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res user.DeleteResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Errorf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}
}

// テスト用のリクエストを実行
func serveHTTP(method string, url string, body io.Reader) *httptest.ResponseRecorder {
	router := myrouter.NewMyRouterWithRoutes(user.Routes)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest(method, url, body)
	router.ServeHTTP(writer, request)
	return writer
}
