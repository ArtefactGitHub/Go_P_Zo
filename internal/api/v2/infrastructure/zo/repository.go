package zo

import (
	"context"
	"database/sql"
	"log"

	derr "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
	infra "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure"
)

type repository struct {
}

func NewRepository() d.Repository {
	return &repository{}
}

func (r *repository) FindAll(ctx context.Context) ([]d.Zo, error) {
	db, err := infra.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryContext(ctx, "SELECT * FROM Zos")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	var z d.Zo
	var result []d.Zo
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

func (r *repository) FindAllByUserId(ctx context.Context, userId int) ([]d.Zo, error) {
	db, err := infra.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryContext(ctx, "SELECT * FROM Zos WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	var z d.Zo
	var result []d.Zo
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

func (r *repository) Find(ctx context.Context, id int) (d.Zo, error) {
	db, err := infra.GetDB(ctx)
	if err != nil {
		return d.Zo{}, err
	}

	z := d.Zo{}
	err = db.QueryRowContext(ctx, "SELECT * FROM Zos WHERE id = ?", id).Scan(
		&z.Id,
		&z.AchievementDate,
		&z.Exp,
		&z.CategoryId,
		&z.Message,
		&z.CreatedAt,
		&z.UpdatedAt,
		&z.UserId)
	if err == sql.ErrNoRows {
		return d.Zo{}, derr.NotFound
	} else if err != nil {
		return d.Zo{}, err
	}

	return z, nil
}

func (r *repository) Create(ctx context.Context, z d.Zo) (int, error) {
	tx, err := infra.GetTX(ctx)
	if err != nil {
		return 0, err
	}

	result, err := tx.ExecContext(ctx, `
		INSERT INTO Zos(id, achievementDate, exp, categoryId, message, createdAt, updatedAt, user_id)
		            values(?, ?, ?, ?, ?, ?, ?, ?)`,
		z.Id,
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

func (r *repository) Update(ctx context.Context, z d.Zo) error {
	tx, err := infra.GetTX(ctx)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
		UPDATE Zos
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

func (r *repository) Delete(ctx context.Context, id int) error {
	tx, err := infra.GetTX(ctx)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
		DELETE FROM Zos
		WHERE id = ?`,
		id)
	if err != nil {
		return err
	}

	return nil
}
