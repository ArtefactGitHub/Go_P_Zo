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
	config, err := config.LoadConfig("../configs/config.yml")
	if err != nil {
		panic(err)
	}

	err = mydb.Init(config)
	if err != nil {
		panic(err)
	}
	defer mydb.Finalize()

	handler := middleware.CreateHandler(
		middleware.NewJwtMiddleware(),
		middleware.NewRouterMiddleware(),
	)
	log.Printf("running on %s", address)
	log.Fatal(http.ListenAndServe(address, handler))
}
