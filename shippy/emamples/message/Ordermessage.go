package message

import (
	"fmt"
	"github.com/micro/go-micro/broker"
	"log"
	"time"
)

var (
	topic = "go.micro.web.topic.order"
)

func Pub(name string,price int64,consignmentName string)  {
	msg := &broker.Message{
		Header: map[string]string{
			"name" : fmt.Sprintf("%s",name),
			"consignmentName":fmt.Sprintf("&s",consignmentName),
			"price": fmt.Sprintf("&d",price),
		},
		Body:[]byte(fmt.Sprintf("%s %s",name,time.Now().Unix())),
	}
	if err := broker.Publish(topic,msg); err != nil{
		log.Print("[pub] send a fail message",err)
	}else {
		log.Print("[pub] send a message",string(msg.Body))
	}
}
