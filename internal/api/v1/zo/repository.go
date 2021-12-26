package zo

import (
	"context"
	"database/sql"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
)

type ZoRepository struct {
}

func (r *ZoRepository) Findall(ctx context.Context) ([]Zo, error) {
	rows, err := mydb.Db.QueryContext(ctx, "SELECT * FROM zos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var z Zo
	var result []Zo
	for rows.Next() {
		err := rows.Scan(
			&z.Id,
			&z.AchievementDate,
			&z.Exp,
			&z.CategoryId,
			&z.Message,
			&z.CreatedAt,
			&z.UpdatedAt,
			&z.UserId)
		if err != nil {
			return nil, err
		}

		result = append(result, z)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *ZoRepository) Find(ctx context.Context, id int) (*Zo, error) {
	z := Zo{}
	err := mydb.Db.QueryRowContext(ctx, "SELECT * FROM zos WHERE id = ?", id).Scan(
		&z.Id,
		&z.AchievementDate,
		&z.Exp,
		&z.CategoryId,
		&z.Message,
		&z.CreatedAt,
		&z.UpdatedAt,
		&z.UserId)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &z, nil
}

func (r *ZoRepository) Create(ctx context.Context, z *Zo) (int, error) {
	result, err := mydb.Db.ExecContext(ctx, `
		INSERT INTO zos(id, achievementDate, exp, categoryId, message, createdAt, updatedAt, user_id)
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
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	z.Id = int(id)
	return z.Id, nil
}

func (r *ZoRepository) CreateTx(ctx context.Context, tx *sql.Tx, z *Zo) (int, error) {
	result, err := tx.ExecContext(ctx, `
		INSERT INTO zos(id, achievementDate, exp, categoryId, message, createdAt, updatedAt, user_id)
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
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	z.Id = int(id)
	return z.Id, nil
}

func (r *ZoRepository) Update(ctx context.Context, z *Zo) error {
	_, err := mydb.Db.ExecContext(ctx, `
		UPDATE zos
		SET achievementDate = ?,
		    exp = ?,
		    categoryId = ?,
		    message = ?,
		    updatedAt = ?,
				user_id = ?
		WHERE id = ?`,
		z.AchievementDate,
		z.Exp,
		z.CategoryId,
		z.Message,
		z.UpdatedAt,
		z.UserId,
		z.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ZoRepository) Delete(ctx context.Context, id int) error {
	_, err := mydb.Db.ExecContext(ctx, `
		DELETE FROM zos
		WHERE id = ?`,
		id)
	if err != nil {
		return err
	}

	return nil
}
