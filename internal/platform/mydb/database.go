package mydb

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	infra "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/config"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mycontext"
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
	FinalizeV2(Db)
}

func FinalizeV2(db *sql.DB) {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Println(err)
			return
		}
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
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			log.Println(err)
		}
	}(tx)

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

// トランザクション処理
// 参考：https://go.dev/doc/database/execute-transactions
func TranV2(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	tx, berr := Db.BeginTx(ctx, nil)
	if berr != nil {
		return nil, berr
	}

	defer func(tx *sql.Tx) {
		if r := recover(); r != nil {
			if er := tx.Rollback(); er != nil {
				log.Printf("tx.Rollback has error: %#v \n", er)
			}
			panic(r)
		} else if err != nil {
			if er := tx.Rollback(); er != nil {
				log.Printf("tx.Rollback has error: %#v \n", er)
			}
		}
	}(tx)

	ctx2 := mycontext.NewContext(ctx, infra.KeyTX, tx)
	result, err := f(ctx2)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return result, nil
}
