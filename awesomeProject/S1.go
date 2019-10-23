package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db , err := sql.Open("mysql", "root:root@tcp(localhost:3306)/gotestdb?charset=utf8")
	if err != nil{
		fmt.Println("connect  err = ",err)
	}
	defer db.Close()
	result , err := db.Exec(`INSERT  userinfo (username,departname,create) VALUES (?,?,?)`,"Z","Golang","19-9-7")
	if err != nil{
		fmt.Println("Exec  err = ",err)
	}
	fmt.Println("result = ",result)
}