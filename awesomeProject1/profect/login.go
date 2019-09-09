package main

import (
	"awesomeProject1/profect/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)
//完成登录校验
func alogin(userID string , userPwd string) (err error){
	//链接到服务器端
	conn , err := net.Dial("tcp","localhost:8898")
	if err != nil{
		fmt.Println("net Dial err = ",err)
		return
	}
	defer conn.Close()

	//准备通过conn发送消息给服务器
	//先定义消息主体
	var mes message.Message
	mes.Type = message.LoginMesType  //消息主体的type

	//3、创建一个LoginMes 结构体  即消息主体的Date
	var loginMes message.LoginMes
	loginMes.UserID = userID
	loginMes.UserPwd = userPwd

	//将loignMes序列化  即消息主体的Date
	data , err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json Marshal err =",err)
		return
	}
	// 把data 赋给 mes.data 字段
	mes.Date = string(data)

	//将 mes 进行序列化  即消息主体
	data ,err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json Marshal err =",err)
		return
	}
	//先获取到data长度 ->转成一个表示长度的byte切片(字节类型的切片)
	var pkglen uint32
	pkglen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkglen)
	//发送消息长度
	n ,err := conn.Write(buf[0:4])
	if n != 4 || err != nil{
		fmt.Println("conn write fail ", err)
		return
	}
	//fmt.Println("客户端 发送的消息长度 = %d\n",len(data))
	//发送消息本身
	n , err = conn.Write(data)
	if err != nil{
		fmt.Println("conn write fail ", err)
		return
	}
	//处理服务器返回信息
	return
}