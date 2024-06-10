package model

import "time"

// 定义数据库user表的模型
type User struct {
	Id         int       `gorm:"column:id"`
	UserId     int64     `gorm:"column:user_id"`
	UserName   string    `gorm:"column:username"`
	Password   string    `gorm:"column:password"`
	Email      string    `gorm:"column:email"`
	Gender     int       `gorm:"column:gender"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}

func (t *User) TableName() string {
	return "user"
}
