package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main(){
	//链接到Redis
	conn , err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil{
		fmt.Println("redis Dial err = ",err)
		return
	}
	//2、通过go 向redis写入数据
	_ , err = conn.Do("Set","name","Golang")
	if err != nil{
		fmt.Println("conn Do err = ",err)
		return
	}
	fmt.Println("---------------------------------------")
	r , err2 :=redis.String (conn.Do("Get","name"))
	if err2 != nil{
		fmt.Println("conn Do err = ",err2)
		return
	}
	fmt.Println(r)
}