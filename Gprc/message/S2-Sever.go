package main

import (
	"Gprc/message"
	"errors"
	"net"
	"net/http"
	"net/rpc"
	"time"
)
type OrderService struct {
	
}

func (os *OrderService) GetOrderInfo(request message.OrderRequest, response *message.OrderInfo)  error  {
	orderMap := map[string]message.OrderInfo{
		"201960001":{OrderId:"201960001",OrderName:"cloth",OrderStaus:"Paid"},
		"201960002":{OrderId:"201960002",OrderName:"Food",OrderStaus:"Paid"},
		"201960003":{OrderId:"201960003",OrderName:"shoes",OrderStaus:"UnPaid"},
	}
	current := time.Now().Unix()
	if (request.TimeStamp > current){
		*response = message.OrderInfo{OrderId: "",OrderName:"",OrderStaus:"Abnormal order status"}
	}else {
		result := orderMap[request.OrderId]
		if result.OrderId != "" {
			*response = orderMap[request.OrderId]
		}else {
			return  errors.New("sever error")
		}
	}
	return nil
}

func main() {
	orderService := new(OrderService)
	err := rpc.Register(orderService)
	if err != nil{
		panic(err.Error())
	}
	rpc.HandleHTTP()
	listen , err := net.Listen("tcp",":8881")
	if err != nil {
		panic(err.Error())
	}
	http.Serve(listen,nil)
}