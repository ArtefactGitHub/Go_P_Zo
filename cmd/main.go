package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/zo"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 設定ファイルの取得
	config, err := config.LoadConfig("../configs/config.yml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("config: %v\n", config)

	err = mydb.Init(config)
	if err != nil {
		panic(err)
	}
	defer mydb.Finalize()

	r := MyRouter{}
	r.Routing()
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type MyRouter struct {
}

func (r *MyRouter) Routing() {
	zc := zo.ZoController{}
	http.HandleFunc("/zo", zc.ZoHandler)
	http.HandleFunc("/zo/", zc.ZoHandler)
}
