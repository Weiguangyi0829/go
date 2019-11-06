package model

type Order struct {
	Username      string `gorm:"column:username"`
	Createtime int64  `gorm:"column:createtime"`
	Ordernumber string `gorm:"column:ordernumber"`
	Price       int64  `gorm:"column:price"`
	Status      string `gorm:"column:Status"`
}
