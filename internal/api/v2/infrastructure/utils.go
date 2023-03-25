package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	KeyDB = "DB"
	KeyTX = "TX"
)

func GetDB(ctx context.Context) (*sql.DB, error) {
	v := ctx.Value(KeyDB)
	value, ok := v.(*sql.DB)
	if !ok {
		return nil, fmt.Errorf("DB object is not set")
	}

	return value, nil
}

func GetTX(ctx context.Context) (*sql.Tx, error) {
	v := ctx.Value(KeyTX)
	value, ok := v.(*sql.Tx)
	if !ok {
		return nil, fmt.Errorf("TX object is not set")
	}

	return value, nil
}
