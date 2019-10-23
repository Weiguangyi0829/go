package main

import (
	go_micro_srv_consignment "Micro/shipper/consignment-service/proto/consignment"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/micro/go-micro"
	"io/ioutil"
	"log"
	"os"
)
const (
	Address  = "localhost:8881"
	DEFAULT_INFO_FILE = "consignment.json"
)

//读取json
func parseFile(filename string) (*go_micro_srv_consignment.Consignment,error){
	fmt.Println("123")
	data , err := ioutil.ReadFile(filename)
	if err != nil {
		return nil ,err
	}
	fmt.Println("456")
	var consignment *go_micro_srv_consignment.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil , errors.New("consignment json is err")
	}
	return consignment , nil
}

func main() {
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	// 创建新的客户端
	client := go_micro_srv_consignment.NewShippingServiceClient("greeter", service.Client())

	// 在命令行中指定新的货物信息 json 文件
	infoFile := DEFAULT_INFO_FILE
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}

	// 解析货物信息
	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v", err)
	}
	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v", err)
	}
	// 新货物是否托运成功
	//log.Printf("created: %t", resp.Created)
	for _, c := range resp.Consignments {
		fmt.Println(c)
		log.Printf("%+v", c)
	}
}