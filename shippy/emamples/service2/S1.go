package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"log"
)

var (
	//topic = "go.micro.web.topic.order"
)

//func sub() {
//	_, err := broker.Subscribe(topic, func(p broker.Event) error {
//		fmt.Printf("[sub] Received Body: %s, Header: %s", string(p.Message().Body), p.Message().Header)
//		return nil
//	}, broker.Queue("go.micro.web.topic.order"))
//	if err != nil {
//		fmt.Println(err)
//	}
//}

func main() {
	bk := broker.NewBroker(
		broker.Addrs(fmt.Sprintf("%s:%d", "127.0.0.1", 12312)),
	)
	_, err := bk.Subscribe(topic, func(p broker.Event) error {
		log.Print("[sub]:Received 111 Body: %s,Header:%s", string(p.Message().Body), p.Message().Header)
		return nil
	})
	if err != nil{
		panic(err)
	}
	s := micro.NewService(
		micro.Name("go.micro.book.web.sub"),
		micro.Broker(bk),
	)
	s.Init()
	if err = s.Run() ; err != nil{
		panic(err)
	}
}