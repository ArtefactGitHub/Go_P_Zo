package zo

import (
	"context"
	"database/sql"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
)

var ac, _ = time.Parse(test.TimeLayout, "2021-12-18")

const userId_1 = 1
const userId_2 = 2

var seeds []Zo = []Zo{
	NewZo(1, ac, 100, 0, "test-1", time.Now(), sql.NullTime{}, userId_1),
	NewZo(2, ac, 200, 0, "test-2", time.Now(), sql.NullTime{}, userId_1),
	NewZo(3, ac, 300, 0, "test-3", time.Now(), sql.NullTime{}, userId_1),
	NewZo(4, ac, 300, 0, "test-4", time.Now(), sql.NullTime{}, userId_2)}

func test_seed(ctx context.Context) {
	_, err := mydb.Tran(ctx, func(ctx context.Context, tx *sql.Tx) (interface{}, error) {
		for _, z := range seeds {
			_, err := mydb.Db.ExecContext(
				ctx,
				`INSERT INTO zos(id, achievementDate, exp, categoryId, message, createdAt, updatedAt, user_id)
										values(?, ?, ?, ?, ?, ?, ?, ?)`,
				nil,
				z.AchievementDate,
				z.Exp,
				z.CategoryId,
				z.Message,
				z.CreatedAt,
				z.UpdatedAt,
				z.UserId)
			if err != nil {
				test.Failuer(err)
			}
		}
		return nil, nil
	})

	// Commit the transaction.
	if err != nil {
		test.Failuer(err)
	}
}
