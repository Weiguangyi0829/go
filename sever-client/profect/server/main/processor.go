package main

import (
	"sever-client/profect/common/message"
	"sever-client/profect/method"
	"sever-client/profect/server/process1"
	"fmt"
	"io"
	"net"
)

type Process2 struct {
	Conn net.Conn
}

//ServerProcessMes 函数 ：根据客户端发送消息的种类不同，决定调用哪个函数
func (this *Process2)ServerProcessMes(mes *message.Message) (err error){
	switch mes.Type {
	    case message.LoginMesType:
			//处理登录登录
			up := &process1.UserloginProcess{
				Conn:this.Conn,
			}
			up.ServerProcessLogin(mes)
		case message.RegisterMesType:
			//处理注册
	default:
		fmt.Println("")
	}
	return
}

func (this *Process2)P() (err error) {
	//循环的客户端发送的信息
	for{
		tf := &method.Transfer{
			Conn: this.Conn,
		}
		//这里读取数据包，直接封装成一个readpkg()
		mes , err := tf.Readpkg()//监听链接
		if err != nil{
			if err == io.EOF{
				fmt.Println("客户端退出，服务端也退出")
				return err
			}else {
				fmt.Println("readpkg err = ",err)
				return  err
			}
		}
		//fmt.Println("mes = ",mes)
		err = this.ServerProcessMes(&mes)
		if err != nil{
			return err
		}
	}
}