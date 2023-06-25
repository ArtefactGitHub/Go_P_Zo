package auth_test

import (
	"context"
	"testing"

	"database/sql"

	test "github.com/ArtefactGitHub/Go_P_Zo/internal/test_v2"
)

var (
	DB *sql.DB
	TX *sql.Tx
)

func TestMain(m *testing.M) {
	var teardown func(db *sql.DB)
	DB, teardown = test.SetupV2()
	defer teardown(DB)

	var err error
	TX, err = DB.BeginTx(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	m.Run()
}
