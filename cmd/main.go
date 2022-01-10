package main

import (
	"log"
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/middleware"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"

	_ "github.com/go-sql-driver/mysql"
)

const address = "localhost:8000"

func main() {
	// 設定ファイルの取得
	cfg, err := config.LoadConfig("../configs/config.yml")
	config.Cfg = cfg
	if err != nil {
		panic(err)
	}

	err = mydb.Init(cfg)
	if err != nil {
		panic(err)
	}
	defer mydb.Finalize()

	handler, err := CreateHandler(cfg)
	if err != nil {
		panic(err)
	}

	log.Printf("running on %s", address)
	log.Fatal(http.ListenAndServe(address, handler))
}

func CreateHandler(config *config.Config) (http.Handler, error) {
	jwt, err := middleware.NewJwtMiddleware(config)
	if err != nil {
		return nil, err
	}

	handler := middleware.CreateHandler(
		jwt,
		middleware.NewRouterMiddleware(),
	)
	return handler, nil
}
