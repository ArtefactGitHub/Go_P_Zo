package zo

import (
	"database/sql"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
)

type zoRepository struct {
}

func (r *zoRepository) findall() ([]zo, error) {
	rows, err := mydb.Db.Query("SELECT * FROM zos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var z zo
	var result []zo
	for rows.Next() {
		err := rows.Scan(
			&z.Id,
			&z.AchievementDate,
			&z.Exp,
			&z.CategoryId,
			&z.Message,
			&z.CreatedAt,
			&z.UpdatedAt)
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

func (r *zoRepository) find(id int) (*zo, error) {
	zo := zo{}
	err := mydb.Db.QueryRow("SELECT * FROM zos WHERE id = ?", id).Scan(
		&zo.Id,
		&zo.AchievementDate,
		&zo.Exp,
		&zo.CategoryId,
		&zo.Message,
		&zo.CreatedAt,
		&zo.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &zo, nil
}

func (r *zoRepository) create(zo *zo) (int, error) {
	result, err := mydb.Db.Exec(`
		INSERT INTO zos(id, achievementDate, exp, categoryId, message, createdAt, updatedAt)
		            values(?, ?, ?, ?, ?, ?, ?)`,
		nil,
		zo.AchievementDate,
		zo.Exp,
		zo.CategoryId,
		zo.Message,
		zo.CreatedAt,
		zo.UpdatedAt)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	zo.Id = int(id)
	return zo.Id, nil
}

func (r *zoRepository) update(zo *zo) error {
	_, err := mydb.Db.Exec(`
		UPDATE zos
		SET achievementDate = ?,
		    exp = ?,
		    categoryId = ?,
		    message = ?,
		    updatedAt = ?
		WHERE id = ?`,
		zo.AchievementDate,
		zo.Exp,
		zo.CategoryId,
		zo.Message,
		zo.UpdatedAt,
		zo.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *zoRepository) delete(id int) error {
	_, err := mydb.Db.Exec(`
		DELETE FROM zos
		WHERE id = ?`,
		id)
	if err != nil {
		return err
	}

	return nil
}
