package main

import (
	"./common"
	proto "./proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"log"
)

func main() {
	var cfg common.GlobalCfg
	config.LoadFile("./config/db.json")
	config.Scan(&cfg)

	// 使用consul做服务发现
	reg := consul.NewRegistry(func(ops *registry.Options) {
		ops.Addrs = []string{
			cfg.Consul["host"],
		}
	})

	// 定义服务
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("renewal.client"),
	)

	service.Init()

	// 创建新的客户端
	renewalClient := proto.NewRenewalReadAoService("renewalDealAo", service.Client())

	header := new(proto.RequestHeader)
	header.Chn = "xys"
	header.ClientType = "h5"
	header.PartnerId = "123"
	header.Source = "h5"

	req := new(proto.GetRenewalRequest)
	req.Header = header
	req.DealId = 73633250
	req.RealTime = false

	rsp, err := renewalClient.GetRenewalDealInfo(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}

	// 打印响应请求
	fmt.Println(rsp.GetDealInfo())
}
