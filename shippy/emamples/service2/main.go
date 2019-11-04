package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"log"
	O "shippy/emamples/proto2"
	"github.com/iGoogle-ink/gopay"
)
const (
	privateKey  = "23 "
	alipayPublicKey = "aaa"
)
type I interface {
	pay(consignmentid1 string, number string, timeout string, cgprice string) (payUrl string)
}

type P struct {}

type S struct {
	P
}

func bodymap(consignmentid1 string, number string, timeout string, cgprice string) gopay.BodyMap {
	bm := make(gopay.BodyMap)
	bm.Set("subject",consignmentid1)
	bm.Set("out_trade_no",number)
	bm.Set("quit_url","123.56.46.125:6666")
	bm.Set("timeout_express",timeout)
	bm.Set("total_amount",cgprice)
	bm.Set("product_code","QUICK_WAP_WAY")//销售产品码，商家和支付宝签约的产品码
	return bm
}

func (p *P) pay(consignmentid1 string, number string, timeout string, cgprice string) (payUrl string,client *gopay.AliPayClient,bm gopay.BodyMap) {
	client = gopay.NewAliPayClient("2016101600703675",privateKey,false)
	fmt.Println("GoPay Version: ", gopay.Version)
	bm = bodymap(consignmentid1 , number , timeout , cgprice )
	payUrl, err := client.AliPayTradePagePay(bm)
	if err != nil{
		panic(err)
	}
	return payUrl ,client,bm
}

func (s *S) Orderpay(ctx context.Context,request *O.Request,response *O.Response) error {
	log.Print("Received pay.Orderpay request")
	payurl,client,bm := s.P.pay()
	aliRsp, err := client.AliPayTradePay(bm)
	if err != nil{
		return err
	}
	ok, err := gopay.VerifyAliPaySign(alipayPublicKey, aliRsp.SignData, aliRsp.Sign)
	if ok {
		response.Payurl = payurl
	}
	return nil
}

func main()  {
	service := micro.NewService(
			micro.Name("go.micro.api.pay"),
	)
	service.Init()
	O.RegisterPAYHandler(service.Server(),new(S))
	if err := service.Run(); err!= nil{
		log.Fatal(err)
	}
}