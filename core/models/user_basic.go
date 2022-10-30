package models

import "time"

type UserBasic struct {
	Id          int
	Identity    string
	Name        string
	Password    string
	Email       string
	NowVolume   int64     `xorm:"now_volume"`
	TotalVolume int64     `xorm:"total_volume"`
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
	DeletedAt   time.Time `xorm:"deleted"`
}

func (table UserBasic) TableName() string {
	return "user_basic"
}
