package main

import (
	"fmt"

	"github.com/ArtefactGitHub/Go_T_TestDBAccess/internal/config"
	"github.com/ArtefactGitHub/Go_T_TestDBAccess/internal/models"
	"github.com/ArtefactGitHub/Go_T_TestDBAccess/internal/models/zo"
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

	zos, err := zo.FindAll()
	if err != nil {
		panic(err)
	}
	for _, zo := range zos {
		fmt.Println(zo)
	}
}
