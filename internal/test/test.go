package test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	_ "github.com/go-sql-driver/mysql"
)

const TimeLayout = "2006-01-02"

func Setup(t *testing.T) func() {
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
