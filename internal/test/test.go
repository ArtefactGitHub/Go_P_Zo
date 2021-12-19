package test

import (
	"fmt"
	"log"
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

func Run(
	t *testing.T,
	tests map[string]func(t *testing.T),
	before func(),
	after func()) {
	// テスト共通のセットアップ
	teardown := Setup(t)
	t.Cleanup(teardown)

	for name, test := range tests {
		if before != nil {
			before()
		}

		t.Run(name, test)

		if after != nil {
			after()
		}
	}
}

func Failuer(err error) {
	log.Panicf("failuer: %v", err)
}

type TrashScanner struct{}

func (TrashScanner) Scan(interface{}) error {
	return nil
}
