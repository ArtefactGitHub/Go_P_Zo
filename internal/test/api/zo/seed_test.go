package zo_test

import (
	"context"
	"database/sql"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/zo"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
)

var ac, _ = time.Parse(test.TimeLayout, "2021-12-18")
var seeds []zo.Zo = []zo.Zo{
	zo.NewZo(1, ac, 100, 0, "test-1", time.Now(), sql.NullTime{}),
	zo.NewZo(2, ac, 200, 0, "test-2", time.Now(), sql.NullTime{}),
	zo.NewZo(3, ac, 300, 0, "test-3", time.Now(), sql.NullTime{})}

func test_seed(ctx context.Context) {
	tx, err := mydb.Db.BeginTx(ctx, nil)
	if err != nil {
		test.Failuer(err)
	}
	// Defer a rollback in case anything fails.
	// https://go.dev/doc/database/execute-transactions
	defer tx.Rollback()

	for _, z := range seeds {
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
			test.Failuer(err)
		}
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		test.Failuer(err)
	}
}
