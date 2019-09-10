package main

import (
	"fmt"
	"os"
)

var userid string
var userPwd string

func main(){
	//接受用户的选择
	var key int
	//判断是否还继续显示菜单
	var loop = true

	for loop {
		fmt.Println("------------欢迎登陆----------------------")
		fmt.Println("\t\t 1、登录")
		fmt.Println("\t\t 2、注册")
		fmt.Println("\t\t 3、退出")

		fmt.Scanf("%d\n",&key)
		switch key {
			case 1:
				fmt.Println("login")
				loop = false
			case 2:
				fmt.Println("create new account")
				loop = false
			case 3:
				fmt.Println("exit")
				os.Exit(0)
			default:
				fmt.Println("please input again")
		}
	}

	//根据用户输入 显示新的提示
	if key == 1{
		//login
		fmt.Println("Your ID")
		fmt.Scanf("%s\n",&userid)
		fmt.Println("Your Password")
		fmt.Scanf("%s\n",&userPwd)
		fmt.Println("1231313")
		err := alogin(userid,userPwd)
		if err != nil{
			fmt.Println("login fail")
		}else {
			fmt.Println("login sucees")
		}
	}else {

	}
}