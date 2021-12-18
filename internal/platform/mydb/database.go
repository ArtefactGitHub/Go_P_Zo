package mydb

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
)

var Db *sql.DB
var err error

// データベースへ接続
func Init(config *config.Config) error {
	Db, err = sql.Open(config.SqlDriver, fmt.Sprintf("%s:%s@%s(%s)/%s",
		config.User,
		config.Password,
		config.Protocol,
		config.Address,
		config.DataBase))

	if err != nil {
		log.Fatal("Open() Error: ", err)
		return err
	}

	if err = Db.Ping(); err != nil {
		log.Fatal("Ping() Error: ", err)
		return err
	}
	return nil
}

func Finalize() {
	if Db != nil {
		Db.Close()
	}
}
