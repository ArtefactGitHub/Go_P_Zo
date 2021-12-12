package zo

import (
	"github.com/ArtefactGitHub/Go_T_TestDBAccess/internal/models"
)

func findall() ([]Zo, error) {
	rows, err := models.Db.Query("SELECT * FROM zos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var zo Zo
	var result []Zo
	for rows.Next() {
		err := rows.Scan(&zo.Id, &zo.AchievementDate, &zo.Exp, &zo.CategoryId, &zo.Message, &zo.CreatedAt, &zo.UpdatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, zo)
	}

	return result, nil
}
