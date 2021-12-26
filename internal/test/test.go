package test

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	_ "github.com/go-sql-driver/mysql"
)

func Run(
	t *testing.T,
	tests map[string]func(t *testing.T),
	before func(),
	after func(),
	seed func(context.Context)) {

	teardown := testInit(t)
	t.Cleanup(teardown)

	for name, test := range tests {
		beforeAll(before, seed)

		t.Run(name, test)

		afterAll(before)
	}
}

func beforeAll(before func(), seed func(context.Context)) {
	ctx := context.Background()
	truncateAll(ctx)
	if seed != nil {
		seed(ctx)
	}

	if before != nil {
		before()
	}
}

func afterAll(after func()) {
	if after != nil {
		after()
	}
}

func testInit(t *testing.T) func() {
	// 設定ファイルの取得
	_, pwd, _, _ := runtime.Caller(0)
	path := fmt.Sprintf("%s/config.yml", filepath.Dir(pwd))
	config, err := config.LoadConfig(path)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = mydb.Init(config)
	if err != nil {
		mydb.Finalize()
		t.Fatalf(err.Error())
	}

	return mydb.Finalize
}

func truncateAll(ctx context.Context) {
	_, err := mydb.Db.Exec("TRUNCATE zos")
	if err != nil {
		Failuer(err)
	}
	_, err = mydb.Db.Exec("TRUNCATE users")
	if err != nil {
		Failuer(err)
	}
}
