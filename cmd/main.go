package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/session"
	v2auth "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/presentation/auth"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/presentation/client"
	v2session "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/presentation/session"
	v2user "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/presentation/user"
	v2zo "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/presentation/zo"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/middleware"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	_ "github.com/go-sql-driver/mysql"
)

const (
	address         = ":8080"
	defaultRootPath = "../"
	defaultHost     = "localhost"
)

var (
	exitCode int
)

func main() {
	defer func() {
		os.Exit(exitCode)
	}()

	// 設定ファイルの取得
	rootPath := getEnv("Go_P_Zo_ROOT_PATH", defaultRootPath)
	cfg, err := config.LoadConfig(rootPath + "configs/config.yml")
	config.Cfg = cfg
	if err != nil {
		exitByError(fmt.Sprintf("could not to load config: %v", err))
		return
	}

	err = mydb.Init(cfg)
	if err != nil {
		exitByError(fmt.Sprintf("could  not connect to database: %v", err))
		return
	}
	defer mydb.Finalize()

	handler, err := createHandler(cfg)
	if err != nil {
		exitByError(fmt.Sprintf("could not create handler: %v", err))
		return
	}

	host := getEnv("HOST", defaultHost)
	log.Printf("running on %s", host+address)
	if err = http.ListenAndServe(host+address, handler); err != nil {
		exitByError(fmt.Sprintf("failed to ListenAndServe: %v", err))
		return
	}
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
			v2session.Routes,
			v2zo.Routes,
			v2auth.Routes,
			v2user.Routes,
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

func exitByError(msg string) {
	log.Println(msg)
	exitCode = 1
}
