package main

import (
	"sever-client/profect/server/model"
	"fmt"
	"net"
	"time"
)


//处理client
func process(conn net.Conn){
	defer conn.Close()
	//调用总控
	pro := &Process2{
		Conn:conn,
	}
	err := pro.P()
	if err != nil{
		fmt.Println("client and sever are fail , err = ",err)
		return
	}
}

func init()  {
	//when the server starts , initialize redis connection pool
	initPool("localhost:6379",16,0,300*time.Second)
	initUserCurd()
}


//完成对UserCurd初始化
func initUserCurd() {
	//这里的pool 本身就是一个全局的变量  redis.go
	//注意初始化顺序问题
	model.MyUserCurd = model.NewUserCurd(pool)
}

func main()  {
	//
	fmt.Println("sever listen123123123123123....")
	listrn , err := net.Listen("tcp","0.0.0.0:8898")
	defer listrn.Close()
	if err != nil{
		fmt.Println("net listn err = ", err)
		return
	}

	//waiting client ......
	for{
		fmt.Println("waiting client connect.....")
		conn , err := listrn.Accept()
		if err != nil{
			fmt.Println("listen accept err = ",err)
		}

		//链接成功 启动一个协程保持通讯
		go process(conn)
	}
}