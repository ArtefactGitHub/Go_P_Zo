package zo

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

// [GET] /zo のルーティングのテスト
func test_route_getall(t *testing.T) {
	writer := serveHTTP("GET", "/zo", nil)

	want := http.StatusOK
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res getAllResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Fatalf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}

	if len(res.Zos) != len(seeds) {
		t.Errorf("Invalid Zos length. len() = %v, want %d", len(res.Zos), len(seeds))
	}
}

// [GET] /zo/:id のルーティングのテスト
func test_route_get(t *testing.T) {
	writer := serveHTTP("GET", "/zo/1", nil)

	want := http.StatusOK
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res getResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Fatalf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}

	if res.Zo.Id != 1 {
		t.Errorf("Invalid Zo. %v", res.Zo)
	}
}

// [POST] /zo のルーティングのテスト
func test_route_post(t *testing.T) {
	ac, _ := time.Parse(test.TimeLayout, "2021-12-18")
	z := newZo(
		0, ac, 555, 0, "created by test_route_post",
		time.Now(), sql.NullTime{})
	j, _ := json.MarshalIndent(z, "", "\t")

	writer := serveHTTP("POST", "/zo", bytes.NewReader(j))

	want := http.StatusCreated
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res postResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Fatalf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}

	if res.Zo.Message != z.Message {
		t.Errorf("Invalid Zo. %v", res.Zo)
	}
}

// [UPDATE] /zo のルーティングのテスト
func test_route_update(t *testing.T) {
	z := seeds[2]
	z.Message = "updated by test_route_update"
	j, _ := json.MarshalIndent(z, "", "\t")

	writer := serveHTTP("PUT", "/zo/1", bytes.NewReader(j))

	want := http.StatusOK
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res putResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Fatalf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}

	if res.Zo.Message != z.Message {
		t.Errorf("Invalid Zo. %v", res.Zo)
	}
}

// [DELETE] /zo のルーティングのテスト
func test_route_delete(t *testing.T) {
	writer := serveHTTP("DELETE", "/zo/1", nil)

	want := http.StatusOK
	if writer.Code != want {
		t.Fatalf("Response code is %v, want %d", writer.Code, want)
	}

	var res deleteResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Errorf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}
}

func serveHTTP(method string, url string, body io.Reader) *httptest.ResponseRecorder {
	mux := http.NewServeMux()
	Routing(mux)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest(method, url, body)
	mux.ServeHTTP(writer, request)

	return writer
}
