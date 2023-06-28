package zo

type Category interface {
	ID() int
	Name() string
	UserID() int
}

type category struct {
	id     int
	name   string
	userID int
}

func NewCategory(
	id int,
	name string,
	userId int,
) Category {
	return &category{
		id:     id,
		name:   name,
		userID: userId}
}

func (v *category) ID() int {
	return v.id
}
func (v *category) Name() string {
	return v.name
}
func (v *category) UserID() int {
	return v.userID
}
