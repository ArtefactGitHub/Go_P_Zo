package zo

import (
	"database/sql"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/mydb"
)

func findall() ([]Zo, error) {
	rows, err := mydb.Db.Query("SELECT * FROM zos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var zo Zo
	var result []Zo
	for rows.Next() {
		err := rows.Scan(
			&zo.Id,
			&zo.AchievementDate,
			&zo.Exp,
			&zo.CategoryId,
			&zo.Message,
			&zo.CreatedAt,
			&zo.UpdatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, zo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func find(id int) (*Zo, error) {
	zo := Zo{}
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

func create(zo *Zo) (int, error) {
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

func update(zo *Zo) error {
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

func delete(id int) error {
	_, err := mydb.Db.Exec(`
		DELETE FROM zos
		WHERE id = ?`,
		id)
	if err != nil {
		return err
	}

	return nil
}
