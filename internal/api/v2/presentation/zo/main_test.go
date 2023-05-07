package zo_test

import (
	"database/sql"
	"testing"

	test "github.com/ArtefactGitHub/Go_P_Zo/internal/test_v2"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func TestMain(m *testing.M) {
	var teardown func(db *sql.DB)
	DB, teardown = test.SetupV2()
	defer teardown(DB)

	m.Run()
}
