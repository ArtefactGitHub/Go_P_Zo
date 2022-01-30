package user

import (
	"context"
	"testing"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
)

var r_tests = map[string]func(t *testing.T){
	"test_user_rep_findall":     test_user_rep_findall,
	"test_user_rep_find":        test_user_rep_find,
	"test_user_rep_create":      test_user_rep_create,
	"test_user_rep_update":      test_user_rep_update,
	"test_user_rep_delete":      test_user_rep_delete,
	"test_userToken_rep_create": test_userToken_rep_create}

func Test_repository(t *testing.T) {
	test.Run(t, r_tests, nil, nil, test_seed)
}

// FindAll()のテスト
func test_user_rep_findall(t *testing.T) {
	r := UserRepository{}
	users, err := r.FindAll(context.Background())
	if err != nil {
		t.Errorf("FindAll() has error: %v", err)
	}

	want := 3
	if len(users) != want {
		t.Errorf("len() = %d, want %d", len(users), want)
	}
}

// Find()のテスト
func test_user_rep_find(t *testing.T) {
	r := UserRepository{}
	u, err := r.Find(context.Background(), 1)
	if err != nil {
		t.Errorf("Find() has error: %v", err)
	}

	if u.FullName() != "山田 太郎" {
		t.Errorf("FullName() = %s, want %s", u.FullName(), "山田 太郎")
	}

	if u.Email != "taro@gmail.com" {
		t.Errorf("Email = %s, want %s", u.Email, "taro@gmail.com")
	}
}

// Create()のテスト
func test_user_rep_create(t *testing.T) {
	r := UserRepository{}
	u := seedUsers[0]
	u.GivenName = "created by test"
	u.Email = "createdbytest@com"
	id, err := r.Create(context.Background(), &u)
	if err != nil {
		t.Fatalf("Create() has error: %v", err)
	}

	want := "created by test"
	var givenName string
	err = mydb.Db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(
		&test.TrashScanner{},
		&givenName,
		&test.TrashScanner{},
		&test.TrashScanner{},
		&test.TrashScanner{},
		&test.TrashScanner{},
		&test.TrashScanner{})
	if err != nil {
		t.Fatalf("Create() has error: %v", err)
	}

	if givenName != want {
		t.Errorf("givenName = %s, want %s", givenName, want)
	}

	var count int
	err = mydb.Db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		t.Fatalf("Create() has error: %v", err)
	}

	if count != len(seedUsers)+1 {
		t.Errorf("count = %d, want %d", count, len(seedUsers)+1)
	}
}

// Update()のテスト
func test_user_rep_update(t *testing.T) {
	r := UserRepository{}
	u := seedUsers[0]
	u.GivenName = "太郎更新"
	err := r.Update(context.Background(), &u)
	if err != nil {
		t.Fatalf("Update() has error: %v", err)
	}

	var givenName string
	err = mydb.Db.QueryRow("SELECT * FROM users WHERE id = ?", u.Id).Scan(
		&test.TrashScanner{},
		&givenName,
		&test.TrashScanner{},
		&test.TrashScanner{},
		&test.TrashScanner{},
		&test.TrashScanner{},
		&test.TrashScanner{})
	if err != nil {
		t.Fatalf("Update() has error: %v", err)
	}

	want := "太郎更新"
	if u.GivenName != want {
		t.Errorf("u.GivenName = %s, want %s", u.GivenName, want)
	}
}

// Delete()のテスト
func test_user_rep_delete(t *testing.T) {
	r := UserRepository{}
	u := seedUsers[0]
	err := r.Delete(context.Background(), u.Id)
	if err != nil {
		t.Fatalf("Delete() has error: %v", err)
	}

	var count int
	want := 2
	err = mydb.Db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		t.Fatalf("Delete() has error: %v", err)
	}

	if count != want {
		t.Errorf("count = %d, want %d", count, want)
	}
}

// Create()のテスト
func test_userToken_rep_create(t *testing.T) {
	r := UserTokenRepository{}
	ut := seedUserTokens[0]
	ut.Token = "created by test"
	id, err := r.Create(context.Background(), &ut)
	if err != nil {
		t.Fatalf("Create() has error: %v", err)
	}

	want := "created by test"
	var token string
	err = mydb.Db.QueryRow("SELECT * FROM UserTokens WHERE id = ?", id).Scan(
		&test.TrashScanner{},
		&test.TrashScanner{},
		&token,
		&test.TrashScanner{},
		&test.TrashScanner{},
		&test.TrashScanner{})
	if err != nil {
		t.Fatalf("Create() has error: %v", err)
	}

	if token != want {
		t.Errorf("token = %s, want %s", token, want)
	}

	var count int
	err = mydb.Db.QueryRow("SELECT COUNT(*) FROM UserTokens").Scan(&count)
	if err != nil {
		t.Fatalf("Create() has error: %v", err)
	}

	if count != len(seedUserTokens)+1 {
		t.Errorf("count = %d, want %d", count, len(seedUsers)+1)
	}
}
