package main

import (
	"database/sql"
	"fmt"

	"github.com/ArtefactGitHub/Go_T_TestDBAccess/pkg/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 設定ファイルの取得
	config, err := config.LoadConfig("../configs/config.yml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("config: %v\n", config)

	// データベース接続
	db, err := sql.Open(config.SqlDriver, fmt.Sprintf("%s:%s@%s(%s)/%s",
		config.User,
		config.Password,
		config.Protocol,
		config.Address,
		config.DataBase))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println(sql.Drivers())
}
