package process1

import (
	"encoding/json"
	"fmt"
	"net"
	"sever-client/profect/common/message"
	"sever-client/profect/method"
	"sever-client/profect/server/model"
)

type UserloginProcess struct {
	Conn net.Conn
	//表示conn是哪个用户的
	UserId int
}


//serverProcessLogin函数，专门处理登录请求
func (this *UserloginProcess)ServerProcessLogin(mes *message.Message) (err error){
	//先从mes 取出mes.Data  并直接反序列化成loginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Date) , &loginMes)
	if err != nil{
		fmt.Println("json Unmarshal fail err = ",err)
		return
	}
	fmt.Println("mes = ",mes)
	//定义一个状态码的消息
	var resMes message.Message  //消息主体
	resMes.Type = message.LoginResMesType  //消息Type
	//消息Data
	var loginResMes message.LoginResMes
	//判断用户名和密码是否正确
	user, err := model.MyUserCurd.Login(loginMes.UserID,loginMes.UserPwd)
	if err != nil{
		if err == model.ERROR_USER_NOTEXISTS{
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		}else if err == model.ERROR_USER_PWD{
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		}else {
			loginResMes.Code = 505
			loginResMes.Error = "sever fail"
		}
	}else {
		loginResMes.Code = 200
		fmt.Println("user = ",user,"suceess")
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
	tf := &method.Transfer{
		Conn: this.Conn,
	}
	err = tf.Writepkg(data)
	return
}

func (this *UserloginProcess)ServerProcessRegister(mes *message.Message) (err error){
	//先从mes 取出mes.Data  并直接反序列化成RegisterMes
	var RegisterMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Date) , &RegisterMes)
	if err != nil{
		fmt.Println("json Unmarshal fail err = ",err)
		return
	}
	fmt.Println("mes1 = ",mes)
	//定义一个状态码的消息
	var resMes message.Message  //消息主体
	resMes.Type = message.RegisterResMesType  //消息Type
	//消息Data
	var RegisterResMes message.RegisterResMes
	//去数据库认证
	err = model.MyUserCurd.Register(&RegisterMes.User)
	if err != nil {
		fmt.Printf("注册用户, 用户id已经占用 err=%v\n", err)
		RegisterResMes.Code = 100
		//return
	} else {
		//没有错误
		fmt.Printf("注册成功了...")
		RegisterResMes.Code = 200
	}
	//RegisterResMes 反序列化 添加到消息主体
	data , err := json.Marshal(RegisterResMes)
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
	tf := &method.Transfer{
		Conn: this.Conn,
	}
	err = tf.Writepkg(data)
	if err != nil {
		fmt.Println("send msg包 failed, ", err)
		return
	}
	return
}