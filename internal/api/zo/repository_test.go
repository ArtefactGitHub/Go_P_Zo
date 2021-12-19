package zo

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
)

var tests map[string]func(t *testing.T) = map[string]func(t *testing.T){
	"_test_findall": _test_findall}

// findall()のテスト
func _test_findall(t *testing.T) {
	r := zoRepository{}
	zos, err := r.findall()
	if err != nil {
		t.Errorf("findall() has error: %v", err)
	}

	want := 3
	if len(zos) != want {
		t.Errorf("len() = %d, want %d", len(zos), want)
	}
}

// repositoryのテスト本体
func Test_repository(t *testing.T) {
	// テスト共通のセットアップ
	teardown := test.Setup(t)
	t.Cleanup(teardown)

	for name, test := range tests {
		// テスト毎のセットアップ
		before()
		t.Run(name, test)
	}
}

func before() {
	_, err := mydb.Db.Exec("TRUNCATE zos")
	if err != nil {
		failuer(err)
	}

	seed()
}

func seed() {
	ctx := context.Background()
	tx, err := mydb.Db.BeginTx(ctx, nil)
	if err != nil {
		failuer(err)
	}
	// Defer a rollback in case anything fails.
	// https://go.dev/doc/database/execute-transactions
	defer tx.Rollback()

	ac, _ := time.Parse(test.TimeLayout, "2021-12-18")
	zos := []zo{newZo(0, ac, 100, 0, "test-1", time.Now(), sql.NullTime{}),
		newZo(0, ac, 200, 0, "test-2", time.Now(), sql.NullTime{}),
		newZo(0, ac, 300, 0, "test-3", time.Now(), sql.NullTime{})}
	for _, z := range zos {
		_, err := mydb.Db.ExecContext(
			ctx,
			`INSERT INTO zos(id, achievementDate, exp, categoryId, message, createdAt, updatedAt)
		            values(?, ?, ?, ?, ?, ?, ?)`,
			nil,
			z.AchievementDate,
			z.Exp,
			z.CategoryId,
			z.Message,
			z.CreatedAt,
			z.UpdatedAt)
		if err != nil {
			failuer(err)
		}
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		failuer(err)
	}
}

func failuer(err error) {
	log.Panicf("failuer: %v", err)
}
