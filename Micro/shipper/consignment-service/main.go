package main

import (
	go_micro_srv_consignment "Gprc/shipper/consignment-service/proto/consignment"
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/micro/go-micro"
)

const (
	port = ":8881"
)
//// 货轮微服务
//service ShippingService {
//// 托运一批货物
//rpc CreateConsignment (Consignment) returns (Response) {
//}
//}
//// 货轮承运的一批货物
//message Consignment {
//string id = 1;                      // 货物编号
//string description = 2;             // 货物描述
//int32 weight = 3;                   // 货物重量
//repeated Container containers = 4;  // 这批货有哪些集装箱
//string vessel_id = 5;               // 承运的货轮
//}
//// 单个集装箱
//message Container {
//string id = 1;          // 集装箱编号
//string customer_id = 2; // 集装箱所属客户的编号
//string origin = 3;      // 出发地
//string user_id = 4;     // 集装箱所属用户的编号
//}
//// 托运结果
//message Response {
//bool created = 1;            // 托运成功
//Consignment consignment = 2;// 新托运的货物
//}

//仓库接口
type IRepository interface {
	Create(consignment *go_micro_srv_consignment.Consignment)(*go_micro_srv_consignment.Consignment,error)
	GetAll() []*go_micro_srv_consignment.Consignment
}

//存放多批货物的仓库 ，实现了IRepository 接口
type Reository struct {
	consignments []*go_micro_srv_consignment.Consignment
}

func (repo *Reository) Create(consignment *go_micro_srv_consignment.Consignment) (*go_micro_srv_consignment.Consignment, error) {
	repo.consignments = append(repo.consignments,consignment)
	return consignment , nil
}
//获取所有货物
func (repo *Reository) GetAll() []*go_micro_srv_consignment.Consignment  {
	return repo.consignments
}

type service struct {
	repo Reository
}
// service 实现 consignment.pb.go 中的 ShippingServiceServer 接口
// 使 service 作为 gRPC 的服务端
func (s *service) CreateConsignment(ctx context.Context,req *go_micro_srv_consignment.Consignment) (*go_micro_srv_consignment.Response,error){
	consignment , err := s.repo.Create(req)
	if err != nil{
		return nil , err
	}
	response := &go_micro_srv_consignment.Response{Created:true,Consignment:consignment}
	return response , nil
}

func (s *service) GetConsignments(ctx context.Context, resquest *go_micro_srv_consignment.GetRequest) (*go_micro_srv_consignment.Response , error) {
		allconsignment := s.repo.GetAll()
		resp := &go_micro_srv_consignment.Response{Consignments:allconsignment}
		return resp,nil
}

func main() {
	//listener , err := net.Listen("tcp",port)
	//if err != nil{
	//	log.Fatalf("failed to listen %v",err)
	//}
	//log.Printf("listen on %s\n",port)
	//sever := grpc.NewServer()
	//repo := Reository{}
	//go_micro_srv_consignment.RegisterShippingServiceServer(sever,&service{repo})
	//if err := sever.Serve(listener); err !=nil{
	//	log.Fatalf("failed to listen %v",err)
	//}
	//log.Printf("listen on 1 %s\n",port)
	server := micro.NewService(

		)
}