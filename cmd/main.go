package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/controllers"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 設定ファイルの取得
	config, err := config.LoadConfig("../configs/config.yml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("config: %v\n", config)

	err = models.Init(config)
	if err != nil {
		panic(err)
	}
	defer models.Finalize()

	http.HandleFunc("/zo", zoHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func zoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		zc := controllers.ZoController{}
		zc.GetAll()

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "MethodGet")
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "MethodPost")
	case http.MethodPut, http.MethodPatch:
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "MethodPut")
	case http.MethodDelete:
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "MethodDelete")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "StatusMethodNotAllowed")
	}
}
