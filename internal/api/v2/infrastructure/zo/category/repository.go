package category

import (
	"context"
	"database/sql"
	"log"
	"time"

	d "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/domain/zo"
	infra "github.com/ArtefactGitHub/Go_P_Zo/internal/api/v2/infrastructure"
)

type (
	categoryRepository struct {
	}

	categoryRecord struct {
		ID        int          `json:"id"`
		UserID    int          `json:"user_id"`
		Name      string       `json:"name"`
		CreatedAt time.Time    `json:"created_at"`
		UpdatedAt sql.NullTime `json:"updated_at"`
	}
)

func NewCategoryRepository() d.CategoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) Finds(ctx context.Context, userId int) ([]d.Category, error) {
	db, err := infra.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryContext(ctx, "SELECT * FROM zo_categories WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	var rec categoryRecord
	var result []categoryRecord
	for rows.Next() {
		err := rows.Scan(
			&rec.ID,
			&rec.UserID,
			&rec.Name,
			&rec.CreatedAt,
			&rec.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, rec)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return toCategoryModels(result), nil
}

func (r *categoryRepository) Create(ctx context.Context, v d.Category) (int, error) {
	tx, err := infra.GetTX(ctx)
	if err != nil {
		return 0, err
	}

	result, err := tx.ExecContext(ctx, `
		INSERT INTO zo_categories(id, name, created_At, updated_At, user_id)
		            values(?, ?, ?, ?, ?)`,
		v.ID(),
		v.Name(),
		time.Now(),
		nil,
		v.UserID())
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func toCategoryModel(rec categoryRecord) d.Category {
	return d.NewCategory(
		rec.ID,
		rec.Name,
		rec.UserID,
	)
}

func toCategoryModels(rec []categoryRecord) []d.Category {
	result := []d.Category{}
	for _, v := range rec {
		z := toCategoryModel(v)
		result = append(result, z)
	}
	return result
}
