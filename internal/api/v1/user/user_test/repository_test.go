package user_test

import (
	"context"
	"testing"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/user"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
)

var r_tests = map[string]func(t *testing.T){
	"user_test_rep_findall": test_r_findall}

func Test_repository(t *testing.T) {
	test.Run(t, r_tests, nil, nil, test_seed)
}

// FindAll()のテスト
func test_r_findall(t *testing.T) {
	r := user.UserRepository{}
	users, err := r.FindAll(context.Background())
	if err != nil {
		t.Errorf("FindAll() has error: %v", err)
	}

	want := 3
	if len(users) != want {
		t.Errorf("len() = %d, want %d", len(users), want)
	}
}
