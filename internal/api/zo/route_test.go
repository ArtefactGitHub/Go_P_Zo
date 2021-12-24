package zo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
)

var route_tests = map[string]func(t *testing.T){
	"test_route_get": test_route_get}

func Test_route(t *testing.T) {
	test.Run(t, route_tests, nil, nil, test_seed)
}

// [GET] zo/:id のルーティングのテスト
func test_route_get(t *testing.T) {
	Routing()

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/zo/1", nil)
	http.DefaultServeMux.ServeHTTP(writer, request)

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
