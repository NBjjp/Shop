package models

type Role struct {
	Id          int
	Title       string
	Description string
	Status      int
	Addtime     int
}

func (Role) TableName() string {
	return "role"
}
