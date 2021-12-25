package zo_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/zo"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
)

var s_tests = map[string]func(t *testing.T){
	"test_s_getAll": test_s_getall,
	"test_s_get":    test_s_get,
	"test_s_post":   test_s_post,
	"test_s_update": test_s_update,
	"test_s_delete": test_s_delete}

func Test_service(t *testing.T) {
	test.Run(t, s_tests, nil, nil, test_seed)
}

// getAll() のテスト
func test_s_getall(t *testing.T) {
	s := zo.ZoService{}
	zos, err := s.GetAll(context.Background())
	if err != nil {
		t.Errorf("getAll() has error: %v", err)
	}

	want := 3
	if len(zos) != want {
		t.Errorf("len() = %d, want %d", len(zos), want)
	}
}

// get() のテスト
func test_s_get(t *testing.T) {
	s := zo.ZoService{}
	z, err := s.Get(context.Background(), 3)
	if err != nil {
		t.Errorf("get() has error: %v", err)
	}

	want := 300
	if z.Exp != want {
		t.Errorf("exp = %d, want %d", z.Exp, want)
	}
}

// post() のテスト
func test_s_post(t *testing.T) {
	s := zo.ZoService{}
	ac, _ := time.Parse(test.TimeLayout, "2021-12-18")
	userId := 1
	z := zo.NewZo(
		0, ac, 555, 0, "created by test",
		time.Now(), sql.NullTime{}, userId)
	_, err := s.Post(context.Background(), &z)
	if err != nil {
		t.Errorf("post() has error: %v", err)
	}

	want := 555
	if z.Exp != want {
		t.Errorf("exp = %d, want %d", z.Exp, want)
	}

	zos, err := s.Zr.Findall(context.Background())
	if err != nil {
		t.Errorf("post() has error: %v", err)
	}
	want = cap(seeds) + 1
	if cap(zos) != want {
		t.Errorf("cap(zos) = %d, want %d", cap(zos), want)
	}
}

// update() のテスト
func test_s_update(t *testing.T) {
	s := zo.ZoService{}
	z := seeds[2]
	z.Message = "updated by test"
	err := s.Update(context.Background(), &z)
	if err != nil {
		t.Errorf("update() has error: %v", err)
	}

	want := "updated by test"
	if z.Message != want {
		t.Errorf("exp = %s, want %s", z.Message, want)
	}
}

// delete() のテスト
func test_s_delete(t *testing.T) {
	s := zo.ZoService{}
	z := seeds[2]
	err := s.Delete(context.Background(), z.Id)
	if err != nil {
		t.Errorf("delete() has error: %v", err)
	}

	zos, err := s.Zr.Findall(context.Background())
	if err != nil {
		t.Errorf("delete() has error: %v", err)
	}
	want := cap(seeds) - 1
	if cap(zos) != want {
		t.Errorf("cap(zos) = %d, want %d", cap(zos), want)
	}
}
