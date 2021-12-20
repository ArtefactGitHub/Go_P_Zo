package zo

import (
	"database/sql"
	"testing"
	"time"

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
	s := zoService{}
	zos, err := s.getAll()
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
	s := zoService{}
	z, err := s.get(3)
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
	s := zoService{}
	ac, _ := time.Parse(test.TimeLayout, "2021-12-18")
	z := newZo(
		0, ac, 555, 0, "created by test",
		time.Now(), sql.NullTime{})
	_, err := s.post(&z)
	if err != nil {
		t.Errorf("post() has error: %v", err)
	}

	want := 555
	if z.Exp != want {
		t.Errorf("exp = %d, want %d", z.Exp, want)
	}

	zos, err := s.zr.findall()
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
	s := zoService{}
	z := seeds[2]
	z.Message = "updated by test"
	err := s.update(&z)
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
	s := zoService{}
	z := seeds[2]
	err := s.delete(z.Id)
	if err != nil {
		t.Errorf("delete() has error: %v", err)
	}

	zos, err := s.zr.findall()
	if err != nil {
		t.Errorf("delete() has error: %v", err)
	}
	want := cap(seeds) - 1
	if cap(zos) != want {
		t.Errorf("cap(zos) = %d, want %d", cap(zos), want)
	}
}
