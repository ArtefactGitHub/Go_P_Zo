package models

import (
	"database/sql"
	"fmt"

	"github.com/ArtefactGitHub/Go_T_TestDBAccess/internal/config"
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

	return err
}

func Finalize() {
	if Db != nil {
		Db.Close()
	}
}
