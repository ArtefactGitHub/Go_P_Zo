package test_v2

import (
	"fmt"
	"log"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
)

func Setup() func() {
	// 設定ファイルの取得
	cfg, err := loadConfig()
	if err != nil {
		log.Panic(err)
	}

	err = mydb.Init(cfg)
	if err != nil {
		mydb.Finalize()
		log.Panic(err)
	}

	fixtures, err := testfixtures.New(
		testfixtures.Database(mydb.Db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory(fmt.Sprintf("%s/testdata/fixtures", currentDir())),
	)
	if err != nil {
		log.Panic(err)
	}
	if err = fixtures.Load(); err != nil {
		log.Panic(err)
	}

	return mydb.Finalize
}

func loadConfig() (*config.Config, error) {
	// 設定ファイルの取得
	path := fmt.Sprintf("%s/config.yml", currentDir())
	cfg, err := config.LoadConfig(path)
	config.Cfg = cfg
	fmt.Printf("config: %#v \n", cfg)
	return cfg, err
}
