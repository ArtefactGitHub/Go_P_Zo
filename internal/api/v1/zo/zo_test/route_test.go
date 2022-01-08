package zo_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/zo"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
)

var route_tests = map[string]func(t *testing.T){
	"test_route_getall": test_route_getall,
	"test_route_get":    test_route_get,
	"test_route_post":   test_route_post,
	"test_route_update": test_route_update,
	"test_route_delete": test_route_delete}

func Test_route(t *testing.T) {
	test.Run(t, route_tests, nil, nil, test_seed)
}

// [GET] /api/v1/zos のルーティングのテスト
func test_route_getall(t *testing.T) {
	writer := serveHTTP("GET", "/api/v1/zos", nil)

	want := http.StatusOK
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res zo.GetAllResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Fatalf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}

	if len(res.Zos) != len(seeds) {
		t.Errorf("Invalid Zos length. len() = %v, want %d", len(res.Zos), len(seeds))
	}
}

// [GET] /api/v1/zos/:zo_id のルーティングのテスト
func test_route_get(t *testing.T) {
	writer := serveHTTP("GET", "/api/v1/zos/1", nil)

	want := http.StatusOK
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res zo.GetResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Fatalf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}

	if res.Zo.Id != 1 {
		t.Errorf("Invalid Zo. %v", res.Zo)
	}
}

// [POST] /api/v1/zos のルーティングのテスト
func test_route_post(t *testing.T) {
	ac, _ := time.Parse(test.TimeLayout, "2021-12-18")
	userId := 1
	z := zo.NewZo(
		0, ac, 555, 0, "created by test_route_post",
		time.Now(), sql.NullTime{}, userId)
	j, _ := json.MarshalIndent(z, "", "\t")

	writer := serveHTTP("POST", "/api/v1/zos", bytes.NewReader(j))

	want := http.StatusCreated
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res zo.PostResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Fatalf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}

	if res.Zo.Message != z.Message {
		t.Errorf("Invalid Zo. %v", res.Zo)
	}
}

// [UPDATE] /api/v1/zos/:zo_id のルーティングのテスト
func test_route_update(t *testing.T) {
	z := seeds[2]
	z.Message = "updated by test_route_update"
	j, _ := json.MarshalIndent(z, "", "\t")

	writer := serveHTTP("PUT", "/api/v1/zos/1", bytes.NewReader(j))

	want := http.StatusOK
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res zo.PutResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Fatalf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}

	if res.Zo.Message != z.Message {
		t.Errorf("Invalid Zo. %v", res.Zo)
	}
}

// [DELETE] /api/v1/zos/:zo_id のルーティングのテスト
func test_route_delete(t *testing.T) {
	writer := serveHTTP("DELETE", "/api/v1/zos/1", nil)

	want := http.StatusOK
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res zo.DeleteResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Errorf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}
}

// テスト用のリクエストを実行
func serveHTTP(method string, url string, body io.Reader) *httptest.ResponseRecorder {
	router := myrouter.NewMyRouterWithRoutes(zo.Routes)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest(method, url, body)
	router.ServeHTTP(writer, request)
	return writer
}
