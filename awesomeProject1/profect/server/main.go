package main

import (
	"awesomeProject1/profect/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func readpkg(conn net.Conn) (mes message.Message, err error)  {
	//将从客户端接受到是已经序列化好的数据
	//读取client 数据
	buf := make([]byte,8096)
	_, err = conn.Read(buf[:4])
	if  err !=nil{
		fmt.Println("conn Read err = ",err )
		return
	}
	//根据buf[:4](是从客户端接收到的数据)转成一个uint32类型
	var pkglen uint32  //包的长度
	pkglen = binary.BigEndian.Uint32(buf[0:4])
	//此时pkglen还是个链接中的uint32类型
	//根据pkglen 读取消息长度  判断是否掉包
	n1 , err := conn.Read(buf[:pkglen])
	if n1 != int(pkglen) || err != nil{
		fmt.Println("conn Read(pkglen) err = ",err )
		return
	}
	//读取消息
	//将连接中的pgklen反序列化  ->message.Message
	err = json.Unmarshal(buf[:pkglen],&mes)
	if err != nil{
		fmt.Println("json.Unmarshal err =" ,err)
		return
	}
	return
}

//处理client
func process(conn net.Conn){
	defer conn.Close()
	//循环的客户端发送的信息
	for{
		//这里读取数据包，直接封装成一个readpkg()
		mes , err := readpkg(conn)
		if err != nil{
			fmt.Println("readpkg err = ",err)
		}
		fmt.Println("mes = ",mes)
	}
}



func main()  {
	fmt.Println("sever listen....")
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