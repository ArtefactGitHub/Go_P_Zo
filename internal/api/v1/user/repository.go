package user

import (
	"context"
	"fmt"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
)

const tableName = "users"

type UserRepository struct {
}

func (r *UserRepository) FindAll(ctx context.Context) ([]User, error) {
	rows, err := mydb.Db.QueryContext(ctx, fmt.Sprintf("SELECT * FROM %s", tableName))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var u User
	var result []User
	for rows.Next() {
		err := rows.Scan(
			&u.Id,
			&u.GivenName,
			&u.FamilyName,
			&u.GivenName,
			&u.Email,
			&u.UpdatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
