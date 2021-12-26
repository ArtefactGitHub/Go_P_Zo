package user_test

import (
	"context"
	"database/sql"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/api/v1/user"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
)

var seeds []user.User = []user.User{
	user.NewUser(1, "太郎", "山田", "taro@gmail.com", time.Now(), sql.NullTime{}),
	user.NewUser(2, "花子", "佐藤", "hanako@gmail.com", time.Now(), sql.NullTime{}),
	user.NewUser(3, "John", "Doe", "john@gmail.com", time.Now(), sql.NullTime{}),
}

func test_seed(ctx context.Context) {
	_, err := mydb.Tran(ctx, func(ctx context.Context, tx *sql.Tx) (interface{}, error) {
		for _, z := range seeds {
			_, err := mydb.Db.ExecContext(
				ctx,
				`INSERT INTO users(id, given_name, family_name, email, createdAt, updatedAt)
									values(?, ?, ?, ?, ?, ?)`,
				nil,
				z.GivenName,
				z.FamilyName,
				z.Email,
				z.CreatedAt,
				z.UpdatedAt)
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
