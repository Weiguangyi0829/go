package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql","root:root@tcp(127.0.0.1:3306)/go?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	fmt.Println("2")
}

type User struct {
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	//创建时间，时间戳
	create_time int64 `gorm:"column:createtime"`
}



func main() {
	fmt.Println("1")
	u := User{
		Username:"tizi365",
		Password:"123456",
		create_time:time.Now().Unix(),
	}
	//插入一条用户数据
	//下面代码会自动生成SQL语句：INSERT INTO `users` (`username`,`password`,`createtime`) VALUES ('tizi365','123456','1540824823')
	//if err := db.Create(u).Error; err != nil {
	//	fmt.Println("插入失败", err)
	//	return
	//}

	//查询并返回第一条数据
	//定义需要保存数据的struct变量
	u = User{}
	isNotFound := db.Where("username = ?","tizi365").First(&u).RecordNotFound()
	if isNotFound {
		fmt.Println("找不到记录")
		return
	}
	//打印查询的数据
	fmt.Println(u.Username,u.Password)
}