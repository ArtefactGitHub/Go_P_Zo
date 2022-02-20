package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/test"
)

var seedUsers []User = []User{
	NewUser(1, "太郎", "山田", "taro@gmail.com", "password", time.Now(), sql.NullTime{}),
	NewUser(2, "花子", "佐藤", "hanako@gmail.com", "password", time.Now(), sql.NullTime{}),
	NewUser(3, "John", "Doe", "john@gmail.com", "password", time.Now(), sql.NullTime{}),
}
var seedUserTokens []UserToken = []UserToken{
	NewUserToken(1, 1, "token_1", time.Now().Add(1*time.Minute), time.Now(), sql.NullTime{}),
	NewUserToken(2, 2, "token_2", time.Now().Add(1*time.Minute), time.Now(), sql.NullTime{}),
	NewUserToken(3, 3, "token_3", time.Now().Add(1*time.Minute), time.Now(), sql.NullTime{}),
}
var seedUserCategories []UserCategory = []UserCategory{
	*NewUserCategory(0, 0, "テストカテゴリー_1", 0, 1, time.Now(), sql.NullTime{}),
	*NewUserCategory(0, 1, "テストカテゴリー_2", 0, 1, time.Now(), sql.NullTime{}),
}

func test_seed(ctx context.Context) {
	_, err := mydb.Tran(ctx, func(ctx context.Context, tx *sql.Tx) (interface{}, error) {
		for _, u := range seedUsers {
			_, err := mydb.Db.ExecContext(
				ctx,
				`INSERT INTO users(id, given_name, family_name, email, password, createdAt, updatedAt)
									values(?, ?, ?, ?, ?, ?, ?)`,
				nil,
				u.GivenName,
				u.FamilyName,
				u.Email,
				u.Password,
				u.CreatedAt,
				u.UpdatedAt)
			if err != nil {
				test.Failuer(err)
			}
		}
		for _, t := range seedUserTokens {
			_, err := mydb.Db.ExecContext(
				ctx,
				`INSERT INTO UserTokens(id, user_id, token, expiredAt, createdAt, updatedAt)
								VALUES(?,?,?,?,?,?)`,
				nil,
				t.UserId,
				t.Token,
				t.ExpiredAt,
				t.CreatedAt,
				t.UpdatedAt)
			if err != nil {
				test.Failuer(err)
			}
		}
		for _, t := range seedUserCategories {
			_, err := mydb.Db.ExecContext(
				ctx,
				`INSERT INTO UserCategories(id, number, name, color_id, user_id, createdAt, updatedAt)
								VALUES(?,?,?,?,?,?,?)`,
				nil,
				t.Number,
				t.Name,
				t.ColorId,
				t.UserId,
				t.CreatedAt,
				t.UpdatedAt)
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
