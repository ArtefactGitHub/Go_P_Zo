package database

import (
	"database/sql"
	"fmt"

	"github.com/ArtefactGitHub/Go_T_TestDBAccess/internal/config"
)

var Db *sql.DB

// データベースへ接続
func Init(config *config.Config) (err error) {
	Db, err := sql.Open(config.SqlDriver, fmt.Sprintf("%s:%s@%s(%s)/%s",
		config.User,
		config.Password,
		config.Protocol,
		config.Address,
		config.DataBase))
	if err != nil {
		return err
	}
	defer Db.Close()

	return nil
}
