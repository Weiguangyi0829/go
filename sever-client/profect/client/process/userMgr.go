package process

import (
	"fmt"
	"sever-client/profect/common/message"
)

//client 维护得map
var onlineUsers map[int]*message.User = make(map[int]*message.User , 10)

//处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes){

	user , ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		//原来没有
		user = &message.User{
			Id:     notifyUserStatusMes.UserId,
		}
	}
	user.Status = notifyUserStatusMes.Satus
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputonlineUser()
}

//在客户端显示当前在线的的用户
func outputonlineUser(){
	//遍历 onlineUsers
	fmt.Println("当前在线用户")
	for id , _ := range onlineUsers{
		fmt.Println("ID：\t",id)

	}
}