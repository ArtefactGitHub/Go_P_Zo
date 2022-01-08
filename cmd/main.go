package main

import (
	"log"
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myrouter"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/user"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/zo"
	_ "github.com/go-sql-driver/mysql"
)

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

	r := Routing()
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}

func Routing() http.Handler {
	r := myrouter.NewMyRouterWithRoutes(
		zo.Routes,
		user.Routes)
	return r
}
