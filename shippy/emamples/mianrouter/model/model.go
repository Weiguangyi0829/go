package model

type User struct {
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	create_time int64 `gorm:"column:createtime"`
}
