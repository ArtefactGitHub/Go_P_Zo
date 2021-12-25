package zo

import (
	"database/sql"
	"time"
)

type Zo struct {
	Id              int          `json:"id"`
	AchievementDate time.Time    `json:"achievementdate"`
	Exp             int          `json:"exp"`
	CategoryId      int          `json:"categoryid"`
	Message         string       `json:"message"`
	CreatedAt       time.Time    `json:"createdat"`
	UpdatedAt       sql.NullTime `json:"updatedat"`
}

func NewZo(
	id int,
	achievementDate time.Time,
	exp int,
	categoryId int,
	message string,
	createdAt time.Time,
	updatedAt sql.NullTime,
) Zo {
	return Zo{
		Id:              id,
		AchievementDate: achievementDate,
		Exp:             exp,
		CategoryId:      categoryId,
		Message:         message,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt}
}
