package main

import (
	proto "../../proto/renewal"
	"../../common"
	"context"
	"encoding/json"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	// 定义服务
	service := micro.NewService(
		micro.Name("go.micro.srv.getRenewalDealByDealId"),
	)

	service.Init()

	// 创建新的客户端
	renewalClient := proto.NewRenewalDealReadAoService("go.micro.srv.getRenewalDealByDealId", service.Client())

	header := new(proto.RequestHeader)
	header.Chn = "xys"
	header.ClientType = "h5"
	header.PartnerId = "123"
	header.Source = "h5"

	req := new(proto.RenewalDealId)
	req.Header = header
	req.DealId = 73633250
	req.RealTime = false

	rsp, err := renewalClient.GetRenewalDealByDealId(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}

	renewalDealInfo := common.ConvertRenewalDealFromProto(rsp.GetRenewalDealInfo())

	data, _ := json.Marshal(renewalDealInfo)

	// 打印响应请求
	log.Println(string(data))
}
