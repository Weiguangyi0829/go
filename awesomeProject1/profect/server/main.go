package main

import (
	"awesomeProject1/profect/common/message"
	 "awesomeProject1/profect/method"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

//serverProcessLogin函数，专门处理登录请求
func serverProcessLogin(conn net.Conn , mes *message.Message) (err error){
	//先从mes 取出mes.Data  并直接反序列化成loginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Date) , &loginMes)
	if err != nil{
		fmt.Println("json Unmarshal fail err = ",err)
		return
	}
	//定义一个状态码的消息
	var resMes message.Message  //消息主体
	resMes.Type = message.LoginResMesType  //消息Type
	//消息Data
	var loginResMes message.LoginResMes
	//判断用户名和密码是否正确
	if loginMes.UserID == "100"  && loginMes.UserPwd == "123" {
		loginResMes.Code = 200
	}else {
		loginResMes.Code = 500
		loginResMes.Error = "login fail"
	}

	//将loginResMes 序列化 添加到消息主体
	data , err := json.Marshal(loginResMes)
	if err != nil{
		fmt.Println("json Marshal err = ",err)
		return
	}
	//将data 赋给 resMes
	resMes.Date = string(data)
	//对resMes进行序列化
	data , err = json.Marshal(resMes)
	if err != nil{
		fmt.Println("json Marshal err = ",err)
		return
	}
	//发送函数
	err = method.Writepkg(conn,data)
	return
}

//ServerProcessMes 函数 ：根据客户端发送消息的种类不同，决定调用哪个函数
func serverProcessMes(conn net.Conn , mes *message.Message) (err error){
	switch mes.Type {
		case message.LoginMesType:
			//处理登录登录
			serverProcessLogin(conn,mes)
		case message.RegisterMesType:
			//处理注册
		default:
			fmt.Println("")
	}
	return
}

//处理client
func process(conn net.Conn){
	defer conn.Close()
	//循环的客户端发送的信息
	for{
		//这里读取数据包，直接封装成一个readpkg()
		mes , err := method.Readpkg(conn)
		if err != nil{
			if err == io.EOF{
				fmt.Println("客户端退出，服务端也退出")
				return
			}else {
				fmt.Println("readpkg err = ",err)
				return
			}
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