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
		return nil, nil
	})

	// Commit the transaction.
	if err != nil {
		test.Failuer(err)
	}
}
