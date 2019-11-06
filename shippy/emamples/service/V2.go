package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"log"
	"shippy/emamples/message"
	Z "shippy/emamples/model"
	b "shippy/emamples/model/db"
	O "shippy/emamples/proto"
	"strconv"
	"time"
)

type R struct {

}

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
	message.Pub(request.Name,i,request.Consignmentid)
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
	service := micro.NewService(
		micro.Name("go.micro.api.CreateOrder"),
	)
	service.Init()
	O.RegisterORDERHandler(service.Server(),new(R))
	if err := service.Run(); err != nil{
		log.Fatal(err)
	}
}