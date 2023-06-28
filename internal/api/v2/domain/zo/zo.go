package zo

import (
	"database/sql"
	"time"
)

type Zo interface {
	ID() int
	AchievementDate() time.Time
	Exp() int
	CategoryID() int
	Category() Category
	Message() string
	CreatedAt() time.Time
	UpdatedAt() sql.NullTime
	UserID() int
}

type zo struct {
	id              int
	achievementDate time.Time
	exp             int
	categoryID      int
	category        Category
	message         string
	createdAt       time.Time
	updatedAt       sql.NullTime
	userID          int
}

func NewZo(
	id int,
	achievementDate time.Time,
	exp int,
	categoryId int,
	message string,
	createdAt time.Time,
	updatedAt sql.NullTime,
	userId int,
) Zo {
	return &zo{
		id:              id,
		achievementDate: achievementDate,
		exp:             exp,
		categoryID:      categoryId,
		message:         message,
		createdAt:       createdAt,
		updatedAt:       updatedAt,
		userID:          userId}
}

func (v *zo) ID() int {
	return v.id
}
func (v *zo) AchievementDate() time.Time {
	return v.achievementDate
}
func (v *zo) Exp() int {
	return v.exp
}
func (v *zo) CategoryID() int {
	return v.categoryID
}
func (v *zo) Category() Category {
	return v.category
}
func (v *zo) Message() string {
	return v.message
}
func (v *zo) CreatedAt() time.Time {
	return v.createdAt
}
func (v *zo) UpdatedAt() sql.NullTime {
	return v.updatedAt
}
func (v *zo) UserID() int {
	return v.userID
}
