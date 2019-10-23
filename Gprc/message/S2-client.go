package main

import (
	"Gprc/message"
	"fmt"
	"net/rpc"
	"time"
)

func main(){
	client , err := rpc.DialHTTP("tcp",":8881")
	if err != nil{
		panic(err.Error())
	}
	timestamp := time.Now().Unix()
	request := message.OrderRequest{OrderId: "201960001",TimeStamp:timestamp}
	var response *message.OrderInfo
	err = client.Call("OrderService.GetOrderInfo",request,&response)
	if err != nil{
		panic(err.Error())
	}
	fmt.Println(*response)
}