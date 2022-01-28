package user

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
)

var s_tests = map[string]func(t *testing.T){
	"test_user_s_getAll": test_user_s_getall,
	"test_user_s_get":    test_user_s_get,
	"test_user_s_post":   test_user_s_post,
	"test_user_s_update": test_user_s_update,
	"test_user_s_delete": test_user_s_delete}

func Test_service(t *testing.T) {
	test.Run(t, s_tests, nil, nil, test_seed)
}

// GetAll() のテスト
func test_user_s_getall(t *testing.T) {
	s := UserService{}
	users, err := s.GetAll(context.Background())
	if err != nil {
		t.Errorf("GetAll() has error: %v", err)
	}

	want := 3
	if len(users) != want {
		t.Errorf("len() = %d, want %d", len(users), want)
	}
}

// Get() のテスト
func test_user_s_get(t *testing.T) {
	s := UserService{}
	u, err := s.Get(context.Background(), 3)
	if err != nil {
		t.Errorf("get() has error: %v", err)
	}

	want := "Doe John"
	if u.FullName() != want {
		t.Errorf("FullName() = %s, want %s", u.FullName(), want)
	}
}

// Post() のテスト
func test_user_s_post(t *testing.T) {
	s := UserService{}
	u := NewUser(1, "太郎更新", "山田", "createbytest@com", "password", time.Now(), sql.NullTime{})
	_, err := s.Post(context.Background(), &u)
	if err != nil {
		t.Errorf("Post() has error: %v", err)
	}

	if u.GivenName != "太郎更新" {
		t.Errorf("GivenName = %s, want %s", u.GivenName, "太郎更新")
	}

	users, err := s.r.FindAll(context.Background())
	if err != nil {
		t.Errorf("Post() has error: %v", err)
	}
	want := cap(seeds) + 1
	if cap(users) != want {
		t.Errorf("cap(users) = %d, want %d", cap(users), want)
	}
}

// Update() のテスト
func test_user_s_update(t *testing.T) {
	s := UserService{}
	u := seeds[2]
	u.GivenName = "updated by test"
	err := s.Update(context.Background(), &u)
	if err != nil {
		t.Errorf("Update() has error: %v", err)
	}

	want := "updated by test"
	if u.GivenName != want {
		t.Errorf("GivenName = %s, want %s", u.GivenName, want)
	}
}

// Delete() のテスト
func test_user_s_delete(t *testing.T) {
	s := UserService{}
	u := seeds[2]
	err := s.Delete(context.Background(), u.Id)
	if err != nil {
		t.Errorf("Delete() has error: %v", err)
	}

	users, err := s.r.FindAll(context.Background())
	if err != nil {
		t.Errorf("Delete() has error: %v", err)
	}
	want := cap(seeds) - 1
	if cap(users) != want {
		t.Errorf("cap(users) = %d, want %d", cap(users), want)
	}
}
