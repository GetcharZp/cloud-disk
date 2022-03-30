package models

type UserBasic struct {
	Id       int
	Identity string
	Name     string
	Password string
	Email    string
}

func (table UserBasic) TableName() string {
	return "user_basic"
}
