package main

import (
	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/client"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/session"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/user"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/zo"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/middleware"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

const address = ":8000"

var defaultRootPath = "../"

func main() {
	// 設定ファイルの取得
	rootPath := getRootPath()
	cfg, err := config.LoadConfig(rootPath + "configs/config.yml")

	config.Cfg = cfg
	if err != nil {
		panic(err)
	}

	err = mydb.Init(cfg)
	if err != nil {
		panic(err)
	}
	defer mydb.Finalize()

	handler, err := createHandler(cfg)
	if err != nil {
		panic(err)
	}

	log.Printf("running on %s", address)
	log.Fatal(http.ListenAndServe(address, handler))
}

func createHandler(config *config.Config) (http.Handler, error) {
	jwt, err := middleware.NewJwtMiddleware(config)
	if err != nil {
		return nil, err
	}
	header, err := middleware.NewHeaderMiddleware()
	if err != nil {
		return nil, err
	}

	handler := middleware.CreateHandler(
		jwt,
		header,
		middleware.NewRouterMiddleware(
			client.Routes,
			session.Routes,
			zo.Routes,
			user.Routes,
		),
	)
	return handler, nil
}

func getRootPath() string {
	path := os.Getenv("Go_P_Zo_ROOT_PATH")
	if path == "" {
		return defaultRootPath
	}

	return path
}
