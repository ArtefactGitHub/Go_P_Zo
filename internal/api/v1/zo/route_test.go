package zo

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/middleware"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myauth"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mycontext"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
	"github.com/ArtefactGitHub/Go_P_Zo/pkg/common"
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
	userId := userId_1
	writer, err := serveHTTP("GET", "/api/v1/zos", nil, userId)
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

	var wantLen int
	for _, m := range seeds {
		if m.UserId == userId {
			wantLen++
		}
	}
	if len(res.Zos) != wantLen {
		t.Errorf("Invalid Zos length. len() = %v, want %d", len(res.Zos), len(seeds))
	}
}

// [GET] /api/v1/zos/:zo_id のルーティングのテスト
func test_route_get(t *testing.T) {
	userId := userId_1
	writer, err := serveHTTP("GET", fmt.Sprintf("/api/v1/zos/%d", userId), nil, userId)
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

	if res.Zo.Id != 1 {
		t.Errorf("Invalid Zo. %v", res.Zo)
	}
}

// [POST] /api/v1/zos のルーティングのテスト
func test_route_post(t *testing.T) {
	ac, _ := time.Parse(test.TimeLayout, "2021-12-18")
	userId := userId_1
	z := NewZo(
		0, ac, 555, 0, "created by test_route_post",
		time.Now(), sql.NullTime{}, userId)
	j, _ := json.MarshalIndent(z, "", "\t")

	writer, err := serveHTTP("POST", "/api/v1/zos", bytes.NewReader(j), userId)
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

	if res.Zo.Message != z.Message {
		t.Errorf("Invalid Zo. %v", res.Zo)
	}
}

// [UPDATE] /api/v1/zos/:zo_id のルーティングのテスト
func test_route_update(t *testing.T) {
	userId := userId_1
	z := seeds[2]
	z.Message = "updated by test_route_update"
	j, _ := json.MarshalIndent(z, "", "\t")

	writer, err := serveHTTP("PUT", fmt.Sprintf("/api/v1/zos/%d", userId), bytes.NewReader(j), userId)
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

	if res.Zo.Message != z.Message {
		t.Errorf("Invalid Zo. %v", res.Zo)
	}
}

// [DELETE] /api/v1/zos/:zo_id のルーティングのテスト
func test_route_delete(t *testing.T) {
	userId := userId_1
	writer, err := serveHTTP("DELETE", fmt.Sprintf("/api/v1/zos/%d", userId), nil, userId)
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

var MockRoutes map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap) = map[myrouter.RouteKey]func(w http.ResponseWriter, r *http.Request, ps common.QueryMap){
	{Path: "/api/v1/zos", Method: "GET", NeedAuth: false}:           zc.getAll,
	{Path: "/api/v1/zos/:zo_id", Method: "GET", NeedAuth: false}:    zc.get,
	{Path: "/api/v1/zos", Method: "POST", NeedAuth: false}:          zc.post,
	{Path: "/api/v1/zos/:zo_id", Method: "PUT", NeedAuth: false}:    zc.update,
	{Path: "/api/v1/zos/:zo_id", Method: "DELETE", NeedAuth: false}: zc.delete,
}

// テスト用のリクエストを実行
func serveHTTP(method string, url string, body io.Reader, userId int) (*httptest.ResponseRecorder, error) {
	writer := httptest.NewRecorder()

	// ユーザートークンをヘッダーへセットしておく
	userToken, err := myauth.CreateUserTokenJwt(userId, time.Now().AddDate(0, 0, 1))
	if err != nil {
		return nil, err
	}
	ctx := mycontext.NewContext(context.Background(), mycontext.UserTokenKey, userToken)
	request, _ := http.NewRequestWithContext(ctx, method, url, body)

	config, _ := test.LoadConfig()
	jwt, err := middleware.NewJwtMiddleware(config)
	if err != nil {
		return nil, err
	}

	handler := middleware.CreateHandler(
		jwt,
		middleware.NewRouterMiddleware(
			MockRoutes,
		),
	)

	handler.ServeHTTP(writer, request)
	return writer, nil
}
