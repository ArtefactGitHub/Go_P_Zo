package zo

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
)

var route_tests = map[string]func(t *testing.T){
	"test_route_getall": test_route_getall,
	"test_route_get":    test_route_get}

func Test_route(t *testing.T) {
	test.Run(t, route_tests, nil, nil, test_seed)
}

// [GET] /zo のルーティングのテスト
func test_route_getall(t *testing.T) {
	writer := setup("GET", "/zo", nil)

	want := 200
	if writer.Code != want {
		t.Errorf("Response code is %v, want %d", writer.Code, want)
	}
	var res getAllResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Errorf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}
	if len(res.Zos) != len(seeds) {
		t.Errorf("Invalid Zos length. len() = %v, want %d", len(res.Zos), len(seeds))
	}
}

// [GET] /zo/:id のルーティングのテスト
func test_route_get(t *testing.T) {
	writer := setup("GET", "/zo/1", nil)

	want := 200
	if writer.Code != want {
		t.Errorf("Response code is %v, want %d", writer.Code, want)
	}
	var res getResponse
	json.Unmarshal(writer.Body.Bytes(), &res)
	if res.StatusCode != want || res.Error != nil {
		t.Errorf("Invalid Response. StatusCode = %d, Error = %v", res.StatusCode, res.Error)
	}
	if res.Zo.Id != 1 {
		t.Errorf("Invalid Zo. %v", res.Zo)
	}
}

func setup(method string, url string, body io.Reader) *httptest.ResponseRecorder {
	mux := http.NewServeMux()
	Routing(mux)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest(method, url, body)
	mux.ServeHTTP(writer, request)

	return writer
}
