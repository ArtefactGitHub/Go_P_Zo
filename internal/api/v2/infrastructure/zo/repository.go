package zo

import (
	"context"
	"database/sql"
	"log"
	"time"

	derr "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/error"
	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
	infra "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure"
)

type (
	repository struct {
	}

	record struct {
		ID              int          `json:"id"`
		AchievementDate time.Time    `json:"achievementdate"`
		Exp             int          `json:"exp"`
		CategoryID      int          `json:"category_id"`
		Message         string       `json:"message"`
		CreatedAt       time.Time    `json:"createdat"`
		UpdatedAt       sql.NullTime `json:"updatedat"`
		UserID          int          `json:"user_id"`
	}
)

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

	var rec record
	var result []record
	for rows.Next() {
		err := rows.Scan(
			&rec.ID,
			&rec.AchievementDate,
			&rec.Exp,
			&rec.CategoryID,
			&rec.Message,
			&rec.CreatedAt,
			&rec.UpdatedAt,
			&rec.UserID)
		if err != nil {
			return nil, err
		}

		result = append(result, rec)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return toModels(result), nil
}

func (r *repository) Finds(ctx context.Context, userId int) ([]d.Zo, error) {
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

	var rec record
	var result []record
	for rows.Next() {
		err := rows.Scan(
			&rec.ID,
			&rec.AchievementDate,
			&rec.Exp,
			&rec.CategoryID,
			&rec.Message,
			&rec.CreatedAt,
			&rec.UpdatedAt,
			&rec.UserID)
		if err != nil {
			return nil, err
		}

		result = append(result, rec)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return toModels(result), nil
}

func (r *repository) Find(ctx context.Context, id int) (d.Zo, error) {
	db, err := infra.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	var rec record
	err = db.QueryRowContext(ctx, "SELECT * FROM Zos WHERE id = ?", id).Scan(
		&rec.ID,
		&rec.AchievementDate,
		&rec.Exp,
		&rec.CategoryID,
		&rec.Message,
		&rec.CreatedAt,
		&rec.UpdatedAt,
		&rec.UserID)
	if err == sql.ErrNoRows {
		return nil, derr.NotFound
	} else if err != nil {
		return nil, err
	}

	return toModel(rec), nil
}

func (r *repository) Create(ctx context.Context, z d.Zo) (int, error) {
	tx, err := infra.GetTX(ctx)
	if err != nil {
		return 0, err
	}

	result, err := tx.ExecContext(ctx, `
		INSERT INTO Zos(id, achievementDate, exp, categoryId, message, createdAt, updatedAt, user_id)
		            values(?, ?, ?, ?, ?, ?, ?, ?)`,
		z.ID(),
		z.AchievementDate(),
		z.Exp(),
		z.CategoryID(),
		z.Message(),
		z.CreatedAt(),
		z.UpdatedAt(),
		z.UserID())
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
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
		z.AchievementDate(),
		z.Exp(),
		z.CategoryID(),
		z.Message(),
		z.UpdatedAt(),
		z.UserID(),
		z.ID())
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

func toModel(rec record) d.Zo {
	return d.NewZo(
		rec.ID,
		rec.AchievementDate,
		rec.Exp,
		rec.CategoryID,
		rec.Message,
		rec.CreatedAt,
		rec.UpdatedAt,
		rec.UserID,
	)
}

func toModels(rec []record) []d.Zo {
	result := []d.Zo{}
	for _, v := range rec {
		z := toModel(v)
		result = append(result, z)
	}
	return result
}
