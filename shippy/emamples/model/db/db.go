package db

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	Z "shippy/emamples/model"
)

func Initialization() *gorm.DB {
	var err error
	db, err := gorm.Open("mysql","root:root@tcp(127.0.0.1:3306)/go?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		panic(err)
	}
	return db
}

func C(db *gorm.DB, u *Z.Order)  {
	if err := db.Create(u).Error ; err !=nil{
		fmt.Println("插入失败", err)
	}
}

func R(db *gorm.DB,username string) ([]Z.Order,error){
	z := []Z.Order{}
	isNotFound := db.Where("username = ?",username).Find(&z).RecordNotFound()
	if isNotFound {
		return z , errors.New("The user is not exitis")
	}
	return z , nil
}

func D(db *gorm.DB,u *Z.Order)  {
	db.Delete(&u)
}
