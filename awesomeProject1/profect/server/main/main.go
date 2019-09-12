package main

import (
	"fmt"
	"net"
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



func main()  {
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