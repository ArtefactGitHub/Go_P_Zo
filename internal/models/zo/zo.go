package zo

import (
	"database/sql"
	"time"
)

type Zo struct {
	Id              int
	AchievementDate time.Time
	Exp             int
	CategoryId      int
	Message         string
	CreatedAt       time.Time
	UpdatedAt       sql.NullTime
}

func FindAll() ([]Zo, error) {
	return findall()
}
