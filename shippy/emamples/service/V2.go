package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"log"
	"shippy/emamples/message"
	Z "shippy/emamples/model"
	b "shippy/emamples/model/db"
	O "shippy/emamples/proto"
	"strconv"
	"time"
)

var (
	c1 = make(chan string)
	returnTopic = "go.micro.web.topic.pay"
)

type R struct {}

func (r *R) Create(ctx context.Context, request *O.Request,response *O.Response)  error   {
	log.Print("Received OrderCreate.Create request")
	response.Order = &O.Order{
		UserId:               request.UserId,
		Name:                 request.Name,
		Price:                request.Price,
		Createtime:           time.Now().Unix(),
		Status:               "Successfully",
	}
	response.Consignmentid = request.Consignmentid
	response.Status = "Successfully"
	i,_ := strconv.ParseInt(request.Price,10,64)
	db := b.Initialization()
	b.C(db,&Z.Order{
		Username:    request.Name,
		Ordernumber: request.Consignmentid,
		Createtime: response.Order.Createtime,
		Price:      i ,
		Status:      "Successfully",
	})
	message.Puborder(request.Name,i,request.Consignmentid)
	url := <- c1
	response.Payurl = url
	return nil
}

func (r *R) GetAll(ctx context.Context, request *O.Request,response *O.Response)  error {
	name  := request.Name
	fmt.Println(name)
	db := b.Initialization()
	z, err := b.R(db,name)
	if err != nil{
		return err
	}
	var o = []*O.Order{}
	for i := 0;i<len(z);i++ {
		t := &O.Order{
			Name:                 z[i].Username,
			//Price:                z[i].Price,
			Status:               z[i].Status,
			Createtime:           z[i].Createtime,
		}
		o = append(o,t)
	}
	fmt.Println(z)
	response.Status = "Successfully"
	response.Orders = o
	fmt.Println(response)
	return nil
}

func main() {
	bk := broker.NewBroker(
		broker.Addrs(fmt.Sprintf("%s:%d", "127.0.0.1", 12312)),
	)
	_, err := bk.Subscribe(returnTopic, func(p broker.Event) error {
		log.Print("[sub]:Received 111 Body: %s,Header:%s", string(p.Message().Body), p.Message().Header)
		c1 <- string(p.Message().Body)
		return nil
	})
	if err != nil{
		panic(err)
	}
	service := micro.NewService(
		micro.Name("go.micro.api.CreateOrder"),
		micro.Broker(bk),
	)
	service.Init()
	O.RegisterORDERHandler(service.Server(),new(R))
	if err := service.Run(); err != nil{
		log.Fatal(err)
	}
}