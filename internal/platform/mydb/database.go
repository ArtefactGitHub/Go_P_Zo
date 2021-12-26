package mydb

import (
	"context"
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

// トランザクション処理
// 参考：https://go.dev/doc/database/execute-transactions
func Tran(ctx context.Context, f func(ctx context.Context, tx *sql.Tx) (interface{}, error)) (interface{}, error) {
	tx, err := Db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	result, err := f(ctx, tx)
	if err != nil {
		return nil, err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return result, nil
}
