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

func LoadConfig() (*config.Config, error) {
	// 設定ファイルの取得
	_, pwd, _, _ := runtime.Caller(0)
	path := fmt.Sprintf("%s/config.yml", filepath.Dir(pwd))
	cfg, err := config.LoadConfig(path)
	config.Cfg = cfg
	return cfg, err
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
	config, err := LoadConfig()
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
	_, err := mydb.Db.Exec("TRUNCATE Zos")
	if err != nil {
		Failuer(err)
	}
	_, err = mydb.Db.Exec("TRUNCATE Users")
	if err != nil {
		Failuer(err)
	}
	_, err = mydb.Db.Exec("TRUNCATE Clients")
	if err != nil {
		Failuer(err)
	}
	_, err = mydb.Db.Exec("TRUNCATE UserTokens")
	if err != nil {
		Failuer(err)
	}
}
