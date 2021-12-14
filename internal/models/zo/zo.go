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

func Find(id int) (*Zo, error) {
	return find(id)
}

func Create(zo *Zo) (int, error) {
	return create(zo)
}

func (zo *Zo) Update() error {
	return update(zo)
}

func Delete(zo *Zo) error {
	return delete(zo)
}
