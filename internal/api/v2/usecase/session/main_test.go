package session_test

import (
	"testing"

	test "github.com/ArtefactGitHub/Go_P_Zo/internal/test_v2"

	_ "github.com/go-sql-driver/mysql"
)

func TestMain(m *testing.M) {
	teardown := test.Setup()
	defer teardown()

	m.Run()
}
