package main

import (
	"fmt"
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

const (
	address         = ":8080"
	defaultRootPath = "../"
	defaultHost     = "localhost"
)

func main() {
	// 設定ファイルの取得
	rootPath := getEnv("Go_P_Zo_ROOT_PATH", defaultRootPath)
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

	host := getEnv("HOST", defaultHost)
	log.Printf("running on %s", host+address)
	log.Fatal(http.ListenAndServe(host+address, handler))
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

func getEnv(key string, defaultValue string) string {
	if val, isSet := os.LookupEnv(key); !isSet {
		fmt.Printf("env[%s] is empty. default is %s \n", key, defaultValue)
		return defaultValue
	} else {
		fmt.Printf("env[%s] is %s \n", key, val)
		return val
	}
}
